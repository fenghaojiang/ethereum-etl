package db

import (
	"context"
	"database/sql"
)

// Design From github.com/bytebase/bytebase

type DBType string

const (
	// ClickHouse is the database type for CLICKHOUSE.
	ClickHouse DBType = "CLICKHOUSE"
	// MySQL is the database type for MYSQL.
	MySQL DBType = "MYSQL"
	// Postgres is the database type for POSTGRES.
	Postgres DBType = "POSTGRES"
	// Snowflake is the database type for SNOWFLAKE.
	Snowflake DBType = "SNOWFLAKE"
	// SQLite is the database type for SQLite.
	SQLite DBType = "SQLITE"
	// TiDB is the database type for TiDB.
	TiDB DBType = "TIDB"
	// MongoDB is the database type for MongoDB.
	MongoDB DBType = "MONGODB"
	// Spanner is the database type for Spanner.
	Spanner DBType = "SPANNER"
	// UnknownDBType is the database type for UNKNOWN.
	UnknownDBType DBType = "UNKNOWN"
)

type DriverConfig struct {
}

type driverFunc func(DriverConfig) IDriver

// ConnectionConfig is the configuration for connections.
type ConnectionConfig struct {
	Host      string
	Port      string
	Username  string
	Password  string
	Database  string
	TLSConfig TLSConfig
	// ReadOnly is only supported for Postgres at the moment.
	ReadOnly bool
	// StrictUseDb will only set as true if the user gives only a database instead of a whole instance to access.
	StrictUseDb bool
	// SRV is only supported for MongoDB now.
	SRV bool
	// AuthenticationDatabase is only supported for MongoDB now.
	AuthenticationDatabase string
}

type IDriver interface {
	Ping(ctx context.Context) error

	Close(ctx context.Context) error
	Open(ctx context.Context, dbType DBType, connectionConfig ConnectionConfig) (IDriver, error)

	GetConnection(ctx context.Context, database string) (*sql.DB, error)
}
