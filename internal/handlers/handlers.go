package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	golog "github.com/Vladroon22/GoLog"
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

func (h *Handlers) AddNewSong(w http.ResponseWriter, r *http.Request) {
	data := h.srv.Data
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		h.logg.Errorln(err)
		return
	}

	id, err := h.repo.AddNewData(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		h.logg.Errorln(err)
		return
	}
	WriteJSON(w, http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

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
}

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
}

func WriteJSON(w http.ResponseWriter, status int, a interface{}) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(a)
}
