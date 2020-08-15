package internal

import (
	"fmt"
	"io"
	"syscall"
)

// LockPendingByte is the first byte of the lock-byte page. If a write lock
// is held on this byte, then a client holds a PENDING lock on this DB.
var LockPendingByte int64 = 0x40000000

// LockReservedByte is the second byte of the lock-byte page. If a write lock
// if held on this byte, then a client holds a RESERVED lock on this DB.
var LockReservedByte int64 = LockPendingByte + 8

// LockSharedFirst marks the start of the shared byte range in the lock-byte page.
// If a read lock is held on the entirety of this range, then at least one
// client holds a SHARED lock on this DB. If a write lock is hold on the entirety
// of the bytes in this range, then a client holds an EXCLUSIVE lock on this DB.
var LockSharedFirst int64 = LockPendingByte + 16

// LockSharedSize is the size in bytes of the shared byte range
// in the lock-byte page.
var LockSharedSize int64 = 510

type LockType int

const (
	LockTypeUnknown LockType = iota
	LockTypeNoLock
	LockTypeShared
	LockTypeReserved
	LockTypePending
	LockTypeExclusive
)

func (db *DB) Lock(typ LockType) (err error) {
	// TODO: confirm whether we need OS-specific implementations of lock commands
	// TODO: do we want to wait? see SETLKW instead of SETLK

	switch typ {
	case LockTypeShared:
		// To obtain a SHARED lock, we:
		//  1. Obtain a read lock on the pending byte
		//  2. Obtain a read lock on the shared byte range
		//  3. Release the read lock on the pending byte
		//
		// #2 is the primary goal, in that once we have a read lock on the shared byte range, then
		// no other process can acquire a non-shared lock. Acquiring a read lock on the pending byte
		// prevents other processes from acquiring a write lock on the pending byte, which is a key
		// step in acquiring a non-shared locks.

		// #1
		if err := syscall.FcntlFlock(db.fd, syscall.F_SETLK, &syscall.Flock_t{
			Len:    1,
			Start:  LockPendingByte,
			Type:   syscall.F_RDLCK,
			Whence: io.SeekStart,
		}); err != nil {
			return err
		}

		// #3: defer s.t. we always release the pending read lock if we have acquired it.
		defer func() {
			if errRelease := syscall.FcntlFlock(db.fd, syscall.F_SETLK, &syscall.Flock_t{
				Len:    1,
				Start:  LockPendingByte,
				Type:   syscall.F_UNLCK,
				Whence: io.SeekStart,
			}); errRelease != nil {
				err = fmt.Errorf("err: %+v. also failed to release PENDING lock: %+v", err, errRelease)
			}
		}()

		// #2
		if err := syscall.FcntlFlock(db.fd, syscall.F_SETLK, &syscall.Flock_t{
			Len:    LockSharedSize,
			Start:  LockSharedFirst,
			Type:   syscall.F_RDLCK,
			Whence: io.SeekStart,
		}); err != nil {
			return err
		}

	// TODO: case LockTypeReserved:
	// TODO: case LockTypePending:
	// TODO: case LockTypeExclusive:
	default:
		return fmt.Errorf("unsupported lock type: %d", typ)
	}

	return nil
}
