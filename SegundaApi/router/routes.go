package router

import (
	"segundaApi/handler"

	"github.com/gin-gonic/gin"
)

func initizalizeRoutes(router *gin.Engine) {

	v1 := router.Group("/api/v1")
	{
		v1.GET("/opening/:id", handler.ShowOpeningHandler)
		v1.PUT("/opening/:id", handler.UpdateOpeningHandler)
		v1.POST("/opening", handler.CreateOpeningHandler)
		v1.DELETE("/opening/:id", handler.DeleteOpeningsHandler)
		v1.GET("/opening", handler.ListOpeningsHandler)
	}
}
