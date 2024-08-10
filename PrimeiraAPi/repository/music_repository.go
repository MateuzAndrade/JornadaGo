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

func (mc *MusicRepository) GetMusicById(id_music int) (*model.Music, error) {
	query, err := mc.connection.Prepare("SELECT * FROM MUSIC WHERE ID = $1")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var music model.Music

	err = query.QueryRow(id_music).Scan(
		&music.ID,
		&music.Music_name,
		&music.Artist_Name,
		&music.Time_Music_Seconds,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	query.Close()
	return &music, nil
}

func (mc *MusicRepository) CreateMusic(music model.Music) (int, error) {
	var id int
	query, err := mc.connection.Prepare("INSERT INTO music" +
		"(music_name,artist_name,time_music_seconds)" +
		" VALUES ($1,$2,$3) RETURNING ID")
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	err = query.QueryRow(music.Music_name, music.Artist_Name, music.Time_Music_Seconds).Scan(&id)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	query.Close()
	return id, nil

}
