package router

import (
	"github.com/gin-gonic/gin"
	"github.com/heitorlimamorei/portifolio-go-api/handlers"
)

func initializeRoutes(router *gin.Engine) {
	handlers.InitHandlers()

	v1 := router.Group("/api/v1/")
	{
		v1.GET("/opening/:id", handlers.GetOpeningHandler)
		v1.POST("/opening", handlers.CreateOpeningHandler)
		v1.DELETE("/opening/:id", handlers.DeleteOpeningHandler)
		v1.PUT("/opening/:id", handlers.UpdateOpeningHandler)
		v1.GET("/openings", handlers.GetOpeningsHandler)
	}
}
