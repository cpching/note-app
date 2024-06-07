package main

import (
    "github.com/gorilla/mux"
)

func registerRoutes(r *mux.Router) {
    api := r.PathPrefix("/api").Subrouter()
    api.HandleFunc("/notes", createNoteHandler).Methods("POST")
    api.HandleFunc("/notes", getNotesHandler).Methods("GET")
    api.HandleFunc("/notes/{id}", getNoteHandler).Methods("GET")
    api.HandleFunc("/notes/{id}", updateNoteHandler).Methods("PUT")
    api.HandleFunc("/notes/{id}", deleteNoteHandler).Methods("DELETE")
    // r.HandleFunc("/notes", createNoteHandler).Methods("POST")
    // r.HandleFunc("/notes", getNotesHandler).Methods("GET")
    // r.HandleFunc("/notes/{id}", getNoteHandler).Methods("GET")
    // r.HandleFunc("/notes/{id}", updateNoteHandler).Methods("PUT")
    // r.HandleFunc("/notes/{id}", deleteNoteHandler).Methods("DELETE")
}

