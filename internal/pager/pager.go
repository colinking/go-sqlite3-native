package pager

import (
	"fmt"
	"os"
	"sync"

	"github.com/pkg/errors"
	"github.com/segmentio/events/v2"
)

// Page is a fixed-size (db.Header.PageSizeBytes) array of bytes. A page
// is unique from an OS-level page and specific to SQLite. The page is
// the low-level abstraction on top of the file system, exposed by the pager,
// which can be read and modified under serializable ACID semantics.
type Page []byte

// Pager is a cache layer on top of an OS file that supports read and write
// methods which operate under serializable ACID semantics.
type Pager struct {
	Path   string
	Header SQLiteHeader

	currentLock      LockType
	currentLockMutex sync.Mutex

	file *os.File
	fd   uintptr

	pid int32
}

func NewPager(path string) (*Pager, error) {
	var err error
	// TODO: add an optional LRU in-memory cache
	// TODO: add page pinning to detect when we should release the SHARED lock.
	//       without an in-memory cache, we can just do global pinning rather than
	//       per-page pinning.
	p := &Pager{
		Path: path,

		currentLock: LockTypeNoLock,

		pid: int32(os.Getpid()),
	}

	events.Debug("opening SQLite DB at: %s", path)

	p.file, err = os.Open(path)
	if err != nil {
		return &Pager{}, errors.Wrap(err, "opening file")
	}
	p.fd = p.file.Fd()

	if err = p.readHeader(); err != nil {
		return &Pager{}, err
	}

	return p, nil
}

// Get returns the contents of the DB page at the provided index, using 1-indexing
// as is convention for page numbers in SQLite.
func (p *Pager) Get(n int) (Page, error) {
	if n < 1 {
		return nil, fmt.Errorf("invalid page index: %d", n)
	}

	// Verify we hold the SHARED lock, acquiring it if we need to:
	if err := p.Lock(LockTypeShared); err != nil {
		return nil, err
	}

	// Load an in-memory copy of the page from the file.
	page := make(Page, p.Header.PageSizeBytes)
	if n > p.Header.DatabaseSizePages {
		// Return an empty page if requesting a page from beyond the edge
		// of the current DB.
		return page, nil
	}
	offset := n * p.Header.PageSizeBytes
	// TODO: look into zero-copying
	_, err := p.file.ReadAt(page, int64(offset))
	if err != nil {
		return nil, err
	}

	return page, nil
}

func (p *Pager) Close() error {
	p.currentLockMutex.Lock()
	defer p.currentLockMutex.Unlock()

	if p.currentLock > LockTypeNoLock {
		if err := p.Unlock(LockTypeNoLock); err != nil {
			return err
		}
	}

	return p.file.Close()
}
