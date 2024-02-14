package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/heitorlimamorei/portifolio-go-api/schemas"
)

func UpdateOpeningHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	req := CreateOpeningRequest{}

	if id == "" {
		sendError(ctx, http.StatusBadRequest, "id is required")
		return
	}

	opening := schemas.Opening{}

	resp := db.First(&opening, id)

	if resp.Error != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("opening with id: %v not found", id))
		return
	}

	ctx.BindJSON(&req)

	if err := req.Validate(); err != nil {
		logger.ErrorF("Error in update opening: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	opening.Company = req.Company
	opening.Role = req.Role
	opening.Location = req.Location
	opening.Salary = req.Salary
	opening.Romote = req.Romote
	opening.Link = req.Link

	db.Save(&opening)

	sendSucess(ctx, "update-opening", opening)

}
