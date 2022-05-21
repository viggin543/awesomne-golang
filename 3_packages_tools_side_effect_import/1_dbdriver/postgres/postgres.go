package postgres

import (
	"database/sql"
	"database/sql/driver"
	"errors"
)

// PostgresDriver provides our implementation for the sql package.
type PostgresDriver struct{}

// Open provides a connection to the database.
// [goland tip]
func (dr PostgresDriver) Open(string) (driver.Conn, error) {
	return nil, errors.New("Unimplemented")
}

var d *PostgresDriver

//db drivers register themselves with the sql package when their init function is executed at startup
//because the sql package can’t know about the drivers that exist when it’s compiled.
func init() {
	d = new(PostgresDriver)
	sql.Register("postgres", d) // how this would have been done without duck typing ? Java example
}
