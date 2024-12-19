package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/pkg/errors"
	"gitlab.com/wit-id/go-orm-mysql/toolkit/db"
)

// NewMySQLDatabase creates a new MySQL database connection.
func NewMySQLDatabase(opt *db.Option) (*sql.DB, error) {
	// Create MySQL connection string
	connURL := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=true",
		opt.Username,
		opt.Password,
		opt.Host,
		opt.Port,
		opt.DatabaseName)

	// Open the database connection
	db, err := sql.Open("mysql", connURL)
	if err != nil {
		return nil, errors.Wrap(err, "mysql: failed to open connection")
	}

	// Set database connection options
	db.SetMaxIdleConns(opt.ConnectionOption.MaxIdle)
	db.SetConnMaxLifetime(opt.ConnectionOption.MaxLifetime)
	db.SetMaxOpenConns(opt.ConnectionOption.MaxOpen)

	// Verify the connection
	ctx, cancel := context.WithTimeout(context.Background(), opt.ConnectionOption.ConnectTimeout)
	defer cancel()

	// Test the connection
	if err := db.PingContext(ctx); err != nil {
		return nil, errors.Wrap(err, "mysql: failed to ping the database")
	}

	log.Println("successfully connected to mysql", connURL)

	// Keep the connection alive
	go doKeepAliveConnection(db, opt.DatabaseName, opt.KeepAliveCheckInterval)

	return db, nil
}

// NewFakeMySQLDB is a mock for testing MySQL database connections.
func NewFakeMySQLDB() (*sql.DB, error) {
	// Use sqlmock to mock a MySQL database
	db, _, err := sqlmock.New()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	return db, nil
}
