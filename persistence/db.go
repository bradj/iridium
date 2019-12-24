package persistence

import (
	"database/sql"
	"fmt"
	"github.com/bradj/iridium/config"
	_ "github.com/lib/pq"
	"log"
)

var db *sql.DB

// NewDB creates a new database connection pool
func NewDB(c config.TomlConfig) (*sql.DB, error) {
	var err error

	connStr := fmt.Sprintf("user='%s' dbname='%s' sslmode=verify-full password='%s'", c.DB.Username, c.DB.Database, c.DB.Password)

	if db, err = sql.Open("postgres", connStr); err != nil {
		log.Fatal(err)
		return nil, err
	}

	return db, nil
}
