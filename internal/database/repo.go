package database

import (
	"time"

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

func (r *Repo) GetInfo() ([]service.Data, error) {
	rows, err := r.db.sqlDB.Query("SELECT name_group, name_song, genre, album, link, lyrics, createAt FROM library")
	if err != nil {
		r.db.logger.Errorln(err)
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
	return data, nil
}

func (r *Repo) GetSong() {

}

func (r *Repo) DeleteData() {

}

func (r *Repo) ChangeData() {

}

func (r *Repo) AddNewData(d ...string) (int, error) {
	data := r.srv.Data
	data.Group_name = d[0]
	data.Song_title = d[1]
	data.Genre = d[3]
	data.Album = d[4]
	data.Lyrics = d[5]
	data.Link = d[6]
	data.Release_date = time.Now().Format(time.DateTime)

	var id int
	query := "INSERT INTO library (group_name, name_song, genre, album, lyrics, link, release_date) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id"
	if err := r.db.sqlDB.QueryRow(query, data.Group_name, data.Song_title, data.Genre, data.Album, data.Lyrics, data.Link, data.Release_date).Scan(&id); err != nil {
		r.db.logger.Errorln(err)
		return 0, err
	}
	r.db.logger.Infoln("new data added")
	return id, nil
}
