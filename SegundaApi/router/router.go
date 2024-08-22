package router

import (
	"github.com/gin-gonic/gin"
)

func Initizalize() {
	router := gin.Default()

	initizalizeRoutes(router)

	router.Run(":8080")
}
