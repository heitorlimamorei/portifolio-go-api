package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/heitorlimamorei/portifolio-go-api/config"
	"gorm.io/gorm"
)

var (
	logger *config.Logger
	db     *gorm.DB
)

func GetOpeningHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "GET Opening",
	})
}

func UpdateOpeningHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "UPDATE Opening",
	})
}

func InitHandlers() {
	logger = config.GetLogger("handlers")
	db = config.GetSqlite()
}
