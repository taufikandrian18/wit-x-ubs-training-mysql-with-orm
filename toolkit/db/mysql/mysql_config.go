package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
	"gitlab.com/wit-id/go-orm-mysql/toolkit/config"
	"gitlab.com/wit-id/go-orm-mysql/toolkit/db"
)

// MySQL-specific function to create a new connection
func NewFromConfig(cfg config.KVStore, path string) (*sql.DB, error) {
	connOpt := db.DefaultConnectionOption()

	// Read connection options
	if maxIdle := cfg.GetInt(fmt.Sprintf("%s.conn.max-idle", path)); maxIdle > 0 {
		connOpt.MaxIdle = maxIdle
	}

	if maxOpen := cfg.GetInt(fmt.Sprintf("%s.conn.max-open", path)); maxOpen > 0 {
		connOpt.MaxOpen = maxOpen
	}

	if maxLifetime := cfg.GetDuration(fmt.Sprintf("%s.conn.max-lifetime", path)); maxLifetime > 0 {
		connOpt.MaxLifetime = maxLifetime
	}

	if connTimeout := cfg.GetDuration(fmt.Sprintf("%s.conn.timeout", path)); connTimeout > 0 {
		connOpt.ConnectTimeout = connTimeout
	}

	if keepAlive := cfg.GetDuration(fmt.Sprintf("%s.conn.keep-alive-interval", path)); keepAlive > 0 {
		connOpt.KeepAliveCheckInterval = keepAlive
	}

	// Construct MySQL connection string
	dbConfig := mysql.Config{
		User:                 cfg.GetString(fmt.Sprintf("%s.username", path)),
		Passwd:               cfg.GetString(fmt.Sprintf("%s.password", path)),
		Net:                  "tcp",
		Addr:                 fmt.Sprintf("%s:%d", cfg.GetString(fmt.Sprintf("%s.host", path)), cfg.GetInt(fmt.Sprintf("%s.port", path))),
		DBName:               cfg.GetString(fmt.Sprintf("%s.schema", path)),
		AllowNativePasswords: true,
		ParseTime:            true,                                                  // Optional, for automatic time type conversion
		Timeout:              cfg.GetDuration(fmt.Sprintf("%s.conn.timeout", path)), // Connect timeout
	}

	// Build the connection string and open the MySQL connection
	dsn := dbConfig.FormatDSN()
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, errors.Wrap(err, "mysql: failed to open connection")
	}

	// Set connection pool options
	db.SetMaxIdleConns(connOpt.MaxIdle)
	db.SetConnMaxLifetime(connOpt.MaxLifetime)
	db.SetMaxOpenConns(connOpt.MaxOpen)

	// Optional: Set a context timeout to check the connection
	ctx, cancel := context.WithTimeout(context.Background(), connOpt.ConnectTimeout)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		return nil, errors.Wrap(err, "mysql: failed to ping database")
	}

	log.Println("successfully connected to mysql", dbConfig.Addr)

	go doKeepAliveConnection(db, cfg.GetString(fmt.Sprintf("%s.schema", path)), connOpt.KeepAliveCheckInterval)

	return db, nil
}

// MySQL-specific keep-alive check
func doKeepAliveConnection(db *sql.DB, dbName string, interval time.Duration) {
	for {
		rows, err := db.Query("SELECT 1")
		if err != nil {
			log.Printf("ERROR db.doKeepAliveConnection conn=mysql error=%s db_name=%s\n", err, dbName)
			return
		}

		if rows.Err() != nil {
			log.Printf("ERROR db.doKeepAliveConnection conn=mysql error=%s db_name=%s\n", rows.Err(), dbName)
			return
		}

		if rows.Next() {
			var i int
			_ = rows.Scan(&i)
			log.Printf("SUCCESS db.doKeepAliveConnection counter=%d db_name=%s stats=%+v\n", i, dbName, db.Stats())
		}

		_ = rows.Close()
		time.Sleep(interval)
	}
}
