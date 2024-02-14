package handlers

import (
	"github.com/heitorlimamorei/portifolio-go-api/config"
	"gorm.io/gorm"
)

var (
	logger *config.Logger
	db     *gorm.DB
)

func InitHandlers() {
	logger = config.GetLogger("handlers")
	db = config.GetSqlite()
}
