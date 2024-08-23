package router

import (
	"segundaApi/handler"

	"github.com/gin-gonic/gin"
)

func initizalizeRoutes(router *gin.Engine) {

	v1 := router.Group("/api/v1")
	{
		v1.GET("/opening", handler.ShowOpeningHandler)
		v1.PUT("/opening", handle)
		v1.POST("/opening")
		v1.DELETE("/opening")
		v1.GET("/opening")
	}
}
