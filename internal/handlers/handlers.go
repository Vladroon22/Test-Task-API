package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	golog "github.com/Vladroon22/GoLog"
	_ "github.com/Vladroon22/TestTask-API/docs"
	"github.com/Vladroon22/TestTask-API/internal/database"
	"github.com/Vladroon22/TestTask-API/internal/service"

	"github.com/gorilla/mux"
)

type Handlers struct {
	srv  *service.Service
	repo *database.Repo
	logg *golog.Logger
}

func NewHandlers(srv *service.Service, rp *database.Repo, lg *golog.Logger) *Handlers {
	return &Handlers{
		srv:  srv,
		repo: rp,
		logg: lg,
	}
}

// @Summary AddingNewSong
// @Description adding new song in library
// @Accept json
// @Produce json
// @Param input body h.srv.Data true "New song data"
// @Success 201 {string} "OK"
// @Failure 400 {string} "Invalid input data"
// @Failure 500 {string} "Internal server error"
// @Router /addSong [post]

func (h *Handlers) AddNewSong(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	group := vars["group"]
	song := vars["song"]

	data := h.srv.Data
	data.Group_name = group
	data.Song_title = song

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		h.logg.Errorln(err)
		return
	}

	err := h.repo.AddNewData(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		h.logg.Errorln(err)
		return
	}
	WriteJSON(w, http.StatusCreated, "OK")
}

// @Summary Get info from library
// @Description Geet information by num of page
// @Accept json
// @Produce json
// @Param id path int true "Page ID"
// @Success 200 {array} Song
// @Failure 500 {string} "Internal server error"
// @Router /info/{id} [get]

func (h *Handlers) GetInfo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	page, _ := strconv.Atoi(vars["id"])
	data, err := h.repo.GetInfo(page)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		h.logg.Errorln(err)
		return
	}
	WriteJSON(w, http.StatusOK, data)
}

// @Summary Get song
// @Description Retrieve song by title and page ID
// @Accept json
// @Produce json
// @Param id path int true "Page ID"
// @Param input body h.srv.Data true "Song data with title"
// @Success 200 {array} h.srv.Data
// @Failure 400 {string} "Invalid input data"
// @Failure 500 {string} "Internal server error"
// @Router /song/{id} [post]

func (h *Handlers) GetSong(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	page, _ := strconv.Atoi(vars["id"])

	data := h.srv.Data
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		h.logg.Errorln(err)
		return
	}

	song, err := h.repo.GetSong(data.Song_title, page)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		h.logg.Errorln(err)
		return
	}
	WriteJSON(w, http.StatusOK, song)
}

// @Summary Delete song
// @Description Delete a song by title
// @Accept json
// @Produce json
// @Param input body h.srv.Data true "Song data with title"
// @Success 204 "Song successfully deleted"
// @Failure 400 {string} "Invalid input data"
// @Failure 500 {string} "Internal server error"
// @Router /deleteSong [delete]

func (h *Handlers) DeleteSong(w http.ResponseWriter, r *http.Request) {
	data := h.srv.Data
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		h.logg.Errorln(err)
		return
	}

	err := h.repo.DeleteData(data.Song_title)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		h.logg.Errorln(err)
		return
	}
	WriteJSON(w, http.StatusNoContent, "Song successfully deleted")
}

// @Summary Change song
// @Description Update song details by title
// @Accept json
// @Produce json
// @Param input body h.srv.DataToChange true "Data to update the song"
// @Success 200 {string} "Updated song details - OK"
// @Failure 400 {string} "Invalid input data"
// @Failure 500 {string} "Internal server error"
// @Router /changeText [patch]

func (h *Handlers) ChangeSong(w http.ResponseWriter, r *http.Request) {
	data := h.srv.DataToChange
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		h.logg.Errorln(err)
		return
	}

	err := h.repo.ChangeData(data.Song, data.DataToChange)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		h.logg.Errorln(err)
		return
	}
	WriteJSON(w, http.StatusOK, "Updated song details - OK")
}

func WriteJSON(w http.ResponseWriter, status int, a interface{}) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(a)
}
