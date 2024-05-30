package main

import (
    "database/sql"
    "encoding/json"
    // "fmt"
    "log"
    "net/http"
    "path/filepath"
    // "strconv"
    "time"

    _ "github.com/go-sql-driver/mysql"
    "github.com/gorilla/mux"
    "github.com/rs/cors"
)

type Note struct {
    ID          int         `json:"id"`
    Title       string      `json:"title"`
    Content     string      `json:"content"`
    CreatedAt   time.Time   `json:"create_at"`
    ModifiedAt  time.Time   `json:"modified_at"`
}

var db *sql.DB

func initDB()  {
    var err error
    if db, err = sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/note_app"); err != nil {
        log.Fatal(err)
    } 
    if err = db.Ping(); err != nil {
        log.Fatal(err)
    }
}

func createNoteHandler(w http.ResponseWriter, r *http.Request)  {
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

func getNotesHandler(w http.ResponseWriter, r *http.Request)  {
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
        layout := "2006-01-02 15:04:05"
        note.CreatedAt, _ = time.Parse(layout, CreatedAt)
        note.ModifiedAt, _ = time.Parse(layout, ModifiedAt)
        notes = append(notes, note)
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(notes)
}

func main() {
    initDB()

    r := mux.NewRouter()

    r.HandleFunc("/notes", createNoteHandler).Methods("POST")
    r.HandleFunc("/notes", getNotesHandler).Methods("GET")

    // Serve static files
    buildPath := filepath.Join("..", "frontend", "note-app-frontend", "build")
    r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir(buildPath))))

    handler := cors.Default().Handler(r)

    log.Println("Server started at :8080")
    if err := http.ListenAndServe(":8080", handler); err != nil {
        log.Fatal(err)
    }
}


