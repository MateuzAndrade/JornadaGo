package router

import "github.com/gin-gonic/gin"

func Initizalize() {
	printMe()
	api := gin.Default()

	api.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	api.Run() // listen and serve on 0.0.0.0:8080
}
