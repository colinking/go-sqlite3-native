package tree

import "github.com/colinking/go-sqlite3-native/internal/pager"

type Tree struct {
	pager *pager.Pager
}

func NewTree(pager *pager.Pager) *Tree {
	return &Tree{
		pager: pager,
	}
}

func (t *Tree) Header() (pager.SQLiteHeader, error) {
	return t.pager.Header()
}

func (t *Tree) Close() error {
	return t.pager.Close()
}
