package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/heitorlimamorei/portifolio-go-api/schemas"
)

func GetOpeningsHandler(ctx *gin.Context) {
	var openings []schemas.Opening
	db.Find(&openings)
	ctx.JSON(http.StatusOK, openings)
}
