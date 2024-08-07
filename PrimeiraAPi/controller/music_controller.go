package controller

import (
	"go-api/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type musicController struct {
}

func NewMusicController() musicController {
	return musicController{}
}

func (p *musicController) GetMusic(ctx *gin.Context) {

	musics := []model.Music{
		{
			ID:                 1,
			Music_name:         "Bad Boys",
			Artist_Name:        "Bob Marley",
			Time_Music_Seconds: 300,
		},
	}

	ctx.JSON(http.StatusOK, musics)
}
