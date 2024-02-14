package config

import (
	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *Logger
)

func InitDb() error {
	var err error

	db, err = InitSqlite()

	return err
}

func GetSqlite() *gorm.DB {
	return db
}

func GetLogger(prefix string) *Logger {
	looger := NewLogger(prefix)
	return looger
}
