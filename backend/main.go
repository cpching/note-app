package main

import (
    "log"
    "net/http"
    "path/filepath"

    "github.com/gorilla/mux"
    "os"
    "github.com/rs/cors"
)

func enableCors(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
        w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
        if r.Method == "OPTIONS" {
            return
        }
        next.ServeHTTP(w, r)
    })
}


func main() {
    initDB()

    r := mux.NewRouter()
    registerRoutes(r)

    // Serve static files
    buildPath := filepath.Join("..", "frontend", "note-app-frontend", "build")

    // Serve static files from the React build directory
    r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(filepath.Join(buildPath, "static")))))

    // Serve the React index.html file for all non-API requests
    r.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        indexPath := filepath.Join(buildPath, "index.html")
        if _, err := os.Stat(indexPath); os.IsNotExist(err) {
            http.NotFound(w, r)
            return
        }
        http.ServeFile(w, r, indexPath)
    })

    handler := cors.Default().Handler(r)

    log.Println("Server started at :8080")
    if err := http.ListenAndServe(":8080", handler); err != nil {
        log.Fatal(err)
    }
}


