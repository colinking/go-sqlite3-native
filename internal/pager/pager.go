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

func (p Page) String() string {
	table := len(p) > 16
	s := ""
	if table {
		s += "OFFSET\n-------------------------------------------------------------\n"
	}

	for i := 0; i < len(p); i += 16 {
		if table {
			s += fmt.Sprintf("0x%08X  ", i)
		}

		for j := i; j < i+16 && j < len(p); j++ {
			s += fmt.Sprintf("%02X ", p[j])
		}

		if table {
			s += "|\n"
		}
	}

	return s
}

// Pager is a cache layer on top of an OS file that supports read and write
// methods which operate under serializable ACID semantics.
type Pager struct {
	mu sync.Mutex

	// TODO: add an optional LRU in-memory cache

	path        string
	header      *SQLiteHeader
	currentLock LockType
	refCount    int
	file        *os.File
	fd          uintptr
	pid         int32
}

func NewPager(path string) (*Pager, error) {
	events.Debug("opening SQLite DB: path=%s", path)

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

	return p, nil
}

// assertSharedWithMutex will acquire a shared lock on the DB file, if not currently held.
//
// Must be called with the pager mutex held.
func (p *Pager) assertSharedWithMutex() error {
	// If we already hold a shared lock, return early.
	if p.currentLock >= LockTypeShared {
		return nil
	}

	// Otherwise, acquire the shared lock:
	if err := p.lock(LockTypeShared); err != nil {
		return err
	}

	// Since we had to acquire the lock, then another writer may have changed
	// the DB since we last held a shared lock. We can check by reloading the header:
	oldHeader := p.header
	if err := p.readHeader(); err != nil {
		return err
	}

	if oldHeader != nil {
		// If the schema cookie number has changed, then a writer operation has occurred
		// since we last held a shared lock. In that case, we need to invalidate the cache.
		if oldHeader.SchemaCookieNumber != p.header.SchemaCookieNumber {
			// TODO: reset the cache
		}
	}

	return nil
}

// TODO: calling Header() then exiting will have not unlocked
func (p *Pager) Header() (SQLiteHeader, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	return p.headerWithMutex()
}

func (p *Pager) headerWithMutex() (SQLiteHeader, error) {
	// In order to read the header, we need to hold the shared lock for concurrency
	// control. If we didn't hold it when we read the header, then we wouldn't
	// be able to guarantee that another writer isn't currently overwriting the
	// header.
	//
	// If we had to acquire the shared lock in order to read the header, then we'll
	// immediately release it at the end of this call.
	//
	// In acquiring the shared lock, this method will refresh p.header if needed.
	if err := p.assertSharedWithMutex(); err != nil {
		return SQLiteHeader{}, err
	}

	if p.refCount == 0 {
		if err := p.unlock(LockTypeNoLock); err != nil {
			return SQLiteHeader{}, err
		}
	}

	return *p.header, nil
}

// Get returns the contents of the DB page at the provided index, using 1-indexing
// as is convention for page numbers in SQLite.
func (p *Pager) Get(n int) (Page, error) {
	if n < 1 {
		return Page{}, fmt.Errorf("invalid page index: %d", n)
	}

	p.mu.Lock()
	defer p.mu.Unlock()

	if err := p.assertSharedWithMutex(); err != nil {
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

	if p.refCount < 0 {
		return fmt.Errorf("too many pages released")
	}

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

	// Verify we've released all pages, otherwise there are pages we are not releasing
	// which means we aren't releasing the shared lock when we can.
	if p.refCount > 0 {
		return fmt.Errorf("pager closed with non-zero refCount (%d)", p.refCount)
	}

	// By the time the Pager is closed, all pages should have been released.
	if p.currentLock != LockTypeNoLock {
		return fmt.Errorf("pager closed but is still locked (refCount=%d)", p.refCount)
	}

	return p.file.Close()
}
