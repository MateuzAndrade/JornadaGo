package repository

import (
	"database/sql"
	"fmt"
	"go-api/model"
)

type MusicRepository struct {
	connection *sql.DB
}

func NewMusicRepository(coonection *sql.DB) MusicRepository {
	return MusicRepository{
		connection: coonection,
	}
}

func (mc *MusicRepository) GetMusics() ([]model.Music, error) {
	query := "SELECT id,music_name,artist_name,time_music_seconds FROM MUSIC"
	rows, err := mc.connection.Query(query)

	if err != nil {
		fmt.Println(err)
		return []model.Music{}, err
	}

	var musicList []model.Music
	var musicObj model.Music

	for rows.Next() {
		err = rows.Scan(
			&musicObj.ID,
			&musicObj.Music_name,
			&musicObj.Artist_Name,
			&musicObj.Time_Music_Seconds)

		if err != nil {
			fmt.Println(err)
			return []model.Music{}, err
		}

		musicList = append(musicList, musicObj)
	}

	rows.Close()

	return musicList, nil

}
