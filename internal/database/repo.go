package database

import (
	"strings"

	"github.com/Vladroon22/TestTask-API/internal/service"
)

type Repo struct {
	db  *DataBase
	srv *service.Service
}

func NewRepo(db *DataBase, srv *service.Service) *Repo {
	return &Repo{
		db:  db,
		srv: srv,
	}
}

func (r *Repo) GetInfo(page int) ([]service.Data, error) {
	limit := 3
	offset := (page - 1) * limit

	query := "SELECT name_group, name_song, genre, album, link, lyrics, createAt FROM library WHERE id > 0 LIMIT $1 OFFSET $2"
	rows, err := r.db.sqlDB.Query(query, limit, offset)
	if err != nil {
		r.db.logger.Errorln(err)
		return nil, err
	}
	defer rows.Close()

	var data []service.Data
	for rows.Next() {
		var dt service.Data
		err := rows.Scan(&dt.Group_name, &dt.Song_title, &dt.Genre, &dt.Album, &dt.Link, &dt.Lyrics, &dt.Release_date)
		if err != nil {
			r.db.logger.Errorln(err)
			return nil, err
		}
		data = append(data, dt)
	}

	if err := rows.Err(); err != nil {
		r.db.logger.Errorln(err)
		return nil, err
	}
	r.db.logger.Infoln("Info gets without errors")
	return data, nil
}

func (r *Repo) GetSong(song string, page int) ([]string, error) {
	var lyrics string
	query := "SELECT lyrics FROM library WHERE name_song = $1"
	err := r.db.sqlDB.QueryRow(query, song).Scan(&lyrics)
	if err != nil {
		return nil, err
	}

	verses := strings.Split(lyrics, `\n`)

	limit := len(verses)
	offset := (page - 1) * page

	if offset >= limit {
		return []string{}, nil
	}

	end := offset + page
	if end > limit {
		end = limit
	}

	r.db.logger.Infoln("the song's text checked")
	return verses[offset:end], nil
}

func (r *Repo) DeleteData(song string) error {
	query := "DELETE FROM library WHERE name_song = $1"
	_, err := r.db.sqlDB.Exec(query, song)
	if err != nil {
		return err
	}
	r.db.logger.Infoln("the " + "'" + song + "'" + " was succesfully deleted")
	return nil
}

func (r *Repo) ChangeData(song string, dataToChange string) error {
	var lyrics string
	query := "SELECT lyrics FROM library WHERE name_song = $1"
	err := r.db.sqlDB.QueryRow(query, song).Scan(&lyrics)
	if err != nil {
		return err
	}

	newLyrics := lyrics + "\n" + dataToChange

	updateQuery := "UPDATE library SET lyrics = $1 WHERE name_song = $2"
	_, err = r.db.sqlDB.Exec(updateQuery, newLyrics, song)
	if err != nil {
		r.db.logger.Errorln(err)
		return err
	}
	r.db.logger.Infoln("all changes complited")
	return nil
}

func (r *Repo) AddNewData(data *service.Data) error {
	query := "INSERT INTO library (group_name, name_song, genre, album, lyrics, link, release_date) VALUES ($1, $2, $3, $4, $5, $6, $7)"
	if _, err := r.db.sqlDB.Exec(query, data.Group_name, data.Song_title, data.Genre, data.Album, data.Lyrics, data.Link, data.Release_date); err != nil {
		r.db.logger.Errorln(err)
		return err
	}
	r.db.logger.Infoln("new data added")
	return nil
}
