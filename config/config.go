package config

import "gorm.io/gorm"

var (
	db     *gorm.DB
	logger *Logger
)

func Init() error {
	return nil
}

func GetLogger(prefix string) *Logger {
	looger := NewLogger(prefix)
	return looger
}
