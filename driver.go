package internal

import (
	"database/sql"
	"database/sql/driver"
)

func init() {
	sql.Register("sqlite3-native", defaultDriver)
}

type Driver struct {
	// TODO
}

var defaultDriver driver.Driver = &Driver{}

func (d *Driver) Open(name string) (driver.Conn, error) {
	// TODO
	return &Conn{}, nil
}
