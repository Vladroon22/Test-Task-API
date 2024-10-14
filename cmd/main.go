package main

import (
	"net/http"
	"os"
	"os/signal"
	"syscall"

	golog "github.com/Vladroon22/GoLog"
	"github.com/Vladroon22/TestTask-API/config"
	"github.com/Vladroon22/TestTask-API/internal/database"
	"github.com/Vladroon22/TestTask-API/internal/handlers"
	"github.com/Vladroon22/TestTask-API/internal/service"
	"github.com/gorilla/mux"
)

func main() {
	cnf := config.CreateConfig()
	logger := golog.New()

	db := database.NewDB(cnf, logger)
	if err := db.Connect(); err != nil {
		logger.Fatalln(err)
		return
	}
	srv := service.NewService()
	repo := database.NewRepo(db, srv)
	h := handlers.NewHandlers(srv, repo, logger)

	router := mux.NewRouter()
	router.HandleFunc("/info/{id:[0-9]+}", h.GetInfo).Methods("GET")
	router.HandleFunc("/song/{id:[0-9]+}", h.GetSong).Methods("GET")
	router.HandleFunc("/addSong", h.AddNewSong).Methods("POST")
	router.HandleFunc("/changeText", h.ChangeSong).Methods("PATCH")
	router.HandleFunc("/deleteSong", h.DeleteSong).Methods("DELETE")

	logger.Infoln("Server is listening --> localhost" + cnf.Addr_PORT)
	go http.ListenAndServe(cnf.Addr_PORT, router)

	exitSig := make(chan os.Signal, 1)
	signal.Notify(exitSig, syscall.SIGINT, syscall.SIGTERM)
	<-exitSig

	go func() {
		if err := db.CloseDB(); err != nil {
			logger.Fatalln(err)
			return
		}
	}()

	logger.Infoln("Gracefull shutdown")
}
