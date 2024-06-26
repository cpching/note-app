package main

import (
	"database/sql"
	"encoding/json"
	// "fmt"

	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

func createNoteHandler(w http.ResponseWriter, r *http.Request) {
	var note Note
	if err := json.NewDecoder(r.Body).Decode(&note); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	query := "INSERT INTO notes (title, content) VALUES (?, ?)"
	result, err := db.Exec(query, note.Title, note.Content)
	if err != nil {
		http.Error(w, "Failed to add note", http.StatusInternalServerError)
		return
	}

	id, _ := result.LastInsertId()
	note.ID = int(id)
	note.CreatedAt = time.Now()
	note.ModifiedAt = time.Now()

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(note)
}

func getNotesHandler(w http.ResponseWriter, r *http.Request) {
	query := "SELECT id, title, content, created_at, modified_at FROM notes"
	rows, err := db.Query(query)
	if err != nil {
		http.Error(w, "Failed to retrieve notes", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	notes := []Note{}
	for rows.Next() {
		var note Note
		var CreatedAt, ModifiedAt string
		if err := rows.Scan(&note.ID, &note.Title, &note.Content, &CreatedAt, &ModifiedAt); err != nil {
			http.Error(w, "Failed to scan note", http.StatusInternalServerError)
			return
		}
		note.CreatedAt, _ = time.Parse(timeLayout, CreatedAt)
		note.ModifiedAt, _ = time.Parse(timeLayout, ModifiedAt)
		notes = append(notes, note)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(notes)
}

func getNoteHandler(w http.ResponseWriter, r *http.Request) {
	// Extract the note ID from the request URL
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid note ID", http.StatusBadRequest)
		return
	}

	var note Note
	var CreatedAt, ModifiedAt string
	query := "SELECT id, title, content, created_at, modified_at FROM notes WHERE id = ?"
	row := db.QueryRow(query, id)
	if err := row.Scan(&note.ID, &note.Title, &note.Content, &CreatedAt, &ModifiedAt); err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Note not fount", http.StatusNotFound)
		} else {
			http.Error(w, "Failed to retrieve note", http.StatusInternalServerError)
		}
		return
	}
	note.CreatedAt, _ = time.Parse(timeLayout, CreatedAt)
	note.ModifiedAt, _ = time.Parse(timeLayout, ModifiedAt)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(note)
}

func updateNoteHandler(w http.ResponseWriter, r *http.Request) {
	// Extract the note ID from the request URL
	params := mux.Vars(r)
	// log.Print(params)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid note ID", http.StatusBadRequest)
		return
	}

	// Decode the JSON request body into a Note struct
	var note Note
	if err := json.NewDecoder(r.Body).Decode(&note); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// log.Println(note)
	modified_at := time.Now()
	// Update the note in the database
	query := "UPDATE notes SET title = ?, content = ?, modified_at = ? WHERE id = ?"
	_, err = db.Exec(query, note.Title, note.Content, modified_at, id)
	if err != nil {
		http.Error(w, "Failed to update note", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(note)
}

func deleteNoteHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid note ID", http.StatusBadRequest)
		return
	}

	query := "DELETE FROM notes WHERE id = ?"
	_, err = db.Exec(query, id)
	if err != nil {
		http.Error(w, "Failed to delete note", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
