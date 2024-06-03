package main

import (
    "log"
    "net/http"
    "path/filepath"

    "github.com/gorilla/mux"
    "github.com/rs/cors"
)

func main() {
    initDB()

    r := mux.NewRouter()
    registerRoutes(r)

    // Serve static files
    buildPath := filepath.Join("..", "frontend", "note-app-frontend", "build")
    r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir(buildPath))))

    handler := cors.Default().Handler(r)

    log.Println("Server started at :8080")
    if err := http.ListenAndServe(":8080", handler); err != nil {
        log.Fatal(err)
    }
}


