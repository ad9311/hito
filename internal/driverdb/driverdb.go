package driverdb

import (
	"database/sql"
	"time"

	"github.com/ad9311/hito/internal/console"

	//
	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

// DB holds the sql database
type DB struct {
	SQL *sql.DB
}

const (
	maxOpenConn = 10
	maxIdleConn = 5
	maxLifeTime = 5 * time.Minute
)

var conn = &DB{}

// ConnectSQL sets up a new database pool for Postgres
func ConnectSQL(dsn string) (*DB, error) {
	db, err := New(dsn)
	console.AssertPanic(err)

	db.SetMaxOpenConns(maxOpenConn)
	db.SetMaxIdleConns(maxIdleConn)
	db.SetConnMaxLifetime(maxLifeTime)

	conn.SQL = db
	err = testDB(db)

	if err != nil {
		return nil, err
	}

	return conn, nil
}

// New creates a new sql connection pool and asserts wether such connection is valid.
func New(dsn string) (*sql.DB, error) {
	console.Log("Opening database connection")
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

func testDB(d *sql.DB) error {
	err := d.Ping()
	if err != nil {
		return err
	}
	return nil
}
