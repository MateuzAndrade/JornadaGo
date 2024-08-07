package main

import (
	"go-api/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	MusicController := controller.NewMusicController()

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Pong Mané",
		})
	})

	server.GET("/musics", MusicController.GetMusic)

	server.Run(":8000")
}
