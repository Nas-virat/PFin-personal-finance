package db

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // The database driver in use.
)

// Config is the database configuration.
type Config struct {
	User       string
	Password   string
	Host       string
	Port       int
	Name       string
	DisableTLS bool
}

// ConnectDB connects to the database.
func ConnectDB(c Config) *sqlx.DB {

	sslmode := "require"

	if c.DisableTLS {
		sslmode = "disable"
	}

	conStr := fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s sslmode=%s", c.User, c.Password, c.Host, c.Port, c.Name, sslmode)

	db, err := sqlx.Open("postgres", conStr)
	if err != nil {
		panic(err)
	}

	db.SetConnMaxLifetime(3 * time.Minute)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	return db
}