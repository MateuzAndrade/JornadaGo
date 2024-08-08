package controller

import (
	"go-api/model"
	"go-api/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type musicController struct {
	MusicUsecase usecase.MusicUsecase
}

func NewMusicController(usecase usecase.MusicUsecase) musicController {
	return musicController{
		MusicUsecase: usecase,
	}
}

func (m *musicController) GetMusic(ctx *gin.Context) {

	musics, err := m.MusicUsecase.GetMusics()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}
	ctx.JSON(http.StatusOK, musics)
}

func (m *musicController) CreateMusic(ctx *gin.Context) {

	var music model.Music
	err := ctx.BindJSON(&music)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	insertMusic, err := m.MusicUsecase.CreateMusic(music)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, insertMusic)
}
