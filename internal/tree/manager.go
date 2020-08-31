package tree

import (
	"github.com/colinking/go-sqlite3-native/internal/pager"
)

type TreeManager struct {
	pager *pager.Pager
}

func NewManager(pager *pager.Pager) *TreeManager {
	return &TreeManager{
		pager: pager,
	}
}

func (tm *TreeManager) Header() (pager.SQLiteHeader, error) {
	return tm.pager.Header()
}

func (tm *TreeManager) Open(rootPage int) (*Tree, error) {
	return newTree(rootPage, tm.pager)
}

func (tm *TreeManager) Close() error {
	return tm.pager.Close()
}
