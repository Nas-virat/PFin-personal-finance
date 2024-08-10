package db

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Config is the database configuration.
type Config struct {
	User       string
	Password   string
	Host       string
	Port       int
	Name       string
	TimeZone   string
	DisableTLS bool
}

// ConnectDB connects to the database.
func ConnectDB(c Config) *gorm.DB {

	sslmode := "require"

	if c.DisableTLS {
		sslmode = "disable"
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s", c.Host, c.User, c.Password, c.Name, c.Port, sslmode, c.TimeZone)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Cann't open gorm postgres %s", err.Error())
	}

	return db
}
