package main

import (
    "github.com/gorilla/mux"
)

func registerRoutes(r *mux.Router) {
    r.HandleFunc("/notes", createNoteHandler).Methods("POST")
    r.HandleFunc("/notes", getNotesHandler).Methods("GET")
    r.HandleFunc("/notes/{id}", updateNoteHandler).Methods("PUT")
}

