package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Vladroon22/TestTask-API/internal/database"
	"github.com/Vladroon22/TestTask-API/internal/service"
)

type Handlers struct {
	srv  *service.Service
	repo *database.Repo
}

func NewHandlers(srv *service.Service, rp *database.Repo) *Handlers {
	return &Handlers{
		srv:  srv,
		repo: rp,
	}
}

func (h *Handlers) AddNewSong(w http.ResponseWriter, r *http.Request) {
	data := h.srv.Data
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func (h *Handlers) GetInfo(w http.ResponseWriter, r *http.Request) {
	data := h.srv.Data
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handlers) GetSong(w http.ResponseWriter, r *http.Request) {

}

func (h *Handlers) DeleteSong(w http.ResponseWriter, r *http.Request) {

}

func (h *Handlers) ChangeSong(w http.ResponseWriter, r *http.Request) {

}
