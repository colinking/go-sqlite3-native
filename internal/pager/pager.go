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
	mu sync.Mutex

	// TODO: add an optional LRU in-memory cache

	path        string
	header      SQLiteHeader
	currentLock LockType
	refCount    int
	file        *os.File
	fd          uintptr
	pid         int32
}

func NewPager(path string) (*Pager, error) {
	events.Debug("opening SQLite DB at: %s", path)

	file, err := os.Open(path)
	if err != nil {
		return &Pager{}, errors.Wrap(err, "opening file")
	}

	p := &Pager{
		path:        path,
		currentLock: LockTypeNoLock,
		refCount:    0,
		file:        file,
		fd:          file.Fd(),
		pid:         int32(os.Getpid()),
	}

	if err := p.readHeader(); err != nil {
		return nil, err
	}

	return p, nil
}

func (p *Pager) assertSharedLocked() error {
	if p.currentLock == LockTypeShared {
		return nil
	}

	// Acquire the shared lock:
	if err := p.lock(LockTypeShared); err != nil {
		return err
	}

	// Now that we have a new shared lock, re-read the header in case it changed:
	oldHeader := p.header
	if err := p.readHeader(); err != nil {
		return err
	}

	if oldHeader.SchemaCookieNumber != p.header.SchemaCookieNumber {
		// TODO: reset the cache
	}

	return nil
}

func (p *Pager) Header() (SQLiteHeader, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	if err := p.assertSharedLocked(); err != nil {
		return SQLiteHeader{}, err
	}

	return p.header, nil
}

// Get returns the contents of the DB page at the provided index, using 1-indexing
// as is convention for page numbers in SQLite.
func (p *Pager) Get(n int) (Page, error) {
	if n < 1 {
		return Page{}, fmt.Errorf("invalid page index: %d", n)
	}

	p.mu.Lock()
	defer p.mu.Unlock()

	if err := p.assertSharedLocked(); err != nil {
		return Page{}, err
	}

	// Load an in-memory copy of the page from the file.
	page := make(Page, p.header.PageSizeBytes)
	// If requesting a page that is beyond the edge of the file, we'll just return
	// an empty page.
	if n <= p.header.DatabaseSizePages {
		offset := (n - 1) * p.header.PageSizeBytes
		_, err := p.file.ReadAt(page, int64(offset))
		if err != nil {
			return Page{}, err
		}
	}

	p.refCount++
	return page, nil
}

func (p *Pager) ReleasePage() error {
	p.mu.Lock()
	defer p.mu.Unlock()

	p.refCount--

	// If all pages have been released, then we can unlock the file.
	if p.refCount == 0 {
		if err := p.unlock(LockTypeNoLock); err != nil {
			return err
		}
	}

	return nil
}

func (p *Pager) Close() error {
	p.mu.Lock()
	defer p.mu.Unlock()

	// Verify we've released all pages, which indicates that we do not hold any locks.
	if p.refCount > 0 {
		return fmt.Errorf("pager closed with non-zero refCount (%d)", p.refCount)
	}

	return p.file.Close()
}
