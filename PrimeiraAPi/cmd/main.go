package main

import (
	"go-api/controller"
	"go-api/db"
	"go-api/repository"
	"go-api/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	//Camada do Repository
	MusicRepository := repository.NewMusicRepository(dbConnection)

	MusicUseCase := usecase.NewMusicUsecase(MusicRepository)

	MusicController := controller.NewMusicController(MusicUseCase)

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Pong Mané",
		})
	})

	server.GET("/musics", MusicController.GetMusic)

	server.Run(":8000")
}
