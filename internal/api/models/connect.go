package models

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB for using GORM
var DB *gorm.DB
var err error

// ConnectDB function
func ConnectDB(dbHost, dbUser, dbPass, dbName, dbPort string) error {
	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	dialect := postgres.Open(dbURL)

	DB, err = gorm.Open(dialect, &gorm.Config{})
	if err != nil {
		return err
	}

	err := Migrate()
	if err != nil {
		return err
	}

	return nil
}
