package usecase

import (
	"go-api/model"
	"go-api/repository"
)

type MusicUsecase struct {
	repository repository.MusicRepository
}

func NewMusicUsecase(repo repository.MusicRepository) MusicUsecase {
	return MusicUsecase{
		repository: repo,
	}
}

func (mu *MusicUsecase) GetMusics() ([]model.Music, error) {
	return mu.repository.GetMusics()
}
