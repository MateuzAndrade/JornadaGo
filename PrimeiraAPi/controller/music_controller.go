package controller

import (
	"go-api/model"
	"go-api/usecase"
	"net/http"
	"strconv"

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

func (m *musicController) GetMusicById(ctx *gin.Context) {

	id := ctx.Param("music_id")

	if id == "" {
		response := model.Response{
			Message: "ID da Music não pode ser Nulo",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	musicId, err := strconv.Atoi(id)
	if err != nil {
		response := model.Response{
			Message: "ID da Music precisa ser um numero",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	music, err := m.MusicUsecase.GetMusicById(musicId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	if music == nil {
		response := model.Response{
			Message: "Music não encontrada",
		}
		ctx.JSON(http.StatusNotFound, response)
		return
	}
	ctx.JSON(http.StatusOK, music)
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
