package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/heitorlimamorei/portifolio-go-api/schemas"
)

func CreateOpeningHandler(ctx *gin.Context) {
	req := CreateOpeningRequest{}

	ctx.BindJSON(&req)

	fmt.Println(req)

	newOpening := schemas.Opening{
		Role:     req.Role,
		Company:  req.Company,
		Location: req.Location,
		Romote:   req.Romote,
		Link:     req.Link,
		Salary:   req.Salary,
	}

	if err := req.Validate(); err != nil {
		logger.ErrorF("Error in creating opening: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	err2 := db.Create(&newOpening).Error

	if err2 != nil {
		logger.ErrorF("Error on saving the opening in the db: %v", err2.Error())
		sendError(ctx, http.StatusInternalServerError, err2.Error())
		return
	}

	sendSucess(ctx, "create-new-opening", newOpening)

}
