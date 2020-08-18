package sqlite3native

import (
	"context"
	"database/sql"
	"database/sql/driver"

	"github.com/colinking/go-sqlite3-native/internal/pager"
	"github.com/colinking/go-sqlite3-native/internal/vm"
)

func init() {
	sql.Register("sqlite3-native", &Driver{})
}

type Driver struct{}

var _ driver.Driver = &Driver{}
var _ driver.DriverContext = &Driver{}

func (d *Driver) Open(name string) (driver.Conn, error) {
	connector, err := d.OpenConnector(name)
	if err != nil {
		return nil, err
	}

	return connector.Connect(context.Background())
}

func (d *Driver) OpenConnector(name string) (driver.Connector, error) {
	// TODO: URI parsing of name to support file:// notation

	return &Connector{
		name:   name,
		driver: d,
	}, nil
}

type Connector struct {
	name string

	driver driver.Driver
}

var _ driver.Connector = &Connector{}

func (c *Connector) Connect(ctx context.Context) (driver.Conn, error) {
	pager, err := pager.NewPager(c.name)
	if err != nil {
		return &Conn{}, err
	}

	m := vm.NewVM()

	return &Conn{
		vm:    m,
		pager: pager,
	}, nil
}

func (c *Connector) Driver() driver.Driver {
	return c.driver
}
