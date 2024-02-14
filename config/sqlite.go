package config

import (
	"os"

	"github.com/heitorlimamorei/portifolio-go-api/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitSqlite() (*gorm.DB, error) {
	dbLogger := GetLogger("sqlite")
	dbPath := "db/main.db"

	// checking if the database is already created
	_, err := os.Stat(dbPath)

	if os.IsNotExist(err) {
		dbLogger.Info("DB file not found, creating it...")

		err := os.MkdirAll("./db", os.ModePerm)

		if err != nil {
			return nil, err
		}

		file, fileError := os.Create(dbPath)

		if fileError != nil {
			return nil, fileError
		}

		file.Close()
	}
	// Opening the database file
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})

	if err != nil {
		dbLogger.ErrorF("Failed to open database: %v", err)
		return nil, err
	}

	// Doing the migrate(that could be create an non existing table or just taking the table) of the Opeing Table
	migrateError := db.AutoMigrate(&schemas.Opening{})

	if migrateError != nil {
		dbLogger.ErrorF("Failed to migrate database: %v", migrateError)
		return nil, migrateError
	}

	return db, nil
}
