package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func NewRouter(db *gorm.DB) http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	}).Methods("GET")

	r.HandleFunc("/tasks", GetTasks(db)).Methods("GET")
	r.HandleFunc("/tasks", CreateTask(db)).Methods("POST")
	r.HandleFunc("/tasks/{id:[0-9]+}", GetTask(db)).Methods("GET")
	r.HandleFunc("/tasks/{id:[0-9]+}", UpdateTask(db)).Methods("PUT")
	r.HandleFunc("/tasks/{id:[0-9]+}", DeleteTask(db)).Methods("DELETE")
	r.HandleFunc("/tasks/{id:[0-9]+}/toggle", ToggleTask(db)).Methods("POST")

	return r
}
