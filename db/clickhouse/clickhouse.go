package clickhouse

import (
	"context"
	"database/sql"

	"github.com/fenghaojiang/ethereum-etl/db"
)

// TODO
type Driver struct {
	db *sql.DB
}

func newDriver(db.DriverConfig) db.IDriver {
	return &Driver{}
}

func (d *Driver) Ping(ctx context.Context) error {
	return nil
}
func (d *Driver) Close(ctx context.Context) error {
	return nil
}
func (d *Driver) Open(ctx context.Context, dbType db.DBType, connectionConfig db.ConnectionConfig) (db.IDriver, error) {
	return nil, nil
}

func (d *Driver) GetConnection(ctx context.Context, database string) (*sql.DB, error) {
	return nil, nil
}

func (d *Driver) Execute(ctx context.Context, statement string) (int64, error) {
	//TODO
	return 0, nil
}
