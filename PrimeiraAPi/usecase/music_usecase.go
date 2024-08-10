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

func (mc *MusicUsecase) GetMusics() ([]model.Music, error) {
	return mc.repository.GetMusics()
}

func (mc *MusicUsecase) GetMusicById(id_music int) (*model.Music, error) {
	music, err := mc.repository.GetMusicById(id_music)
	if err != nil {
		return nil, err
	}
	return music, nil
}

func (mc *MusicUsecase) CreateMusic(music model.Music) (model.Music, error) {

	musicId, err := mc.repository.CreateMusic(music)
	if err != nil {
		return model.Music{}, err
	}

	music.ID = musicId

	return music, nil

}
