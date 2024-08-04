package main

import (
	"log"
	"net/http"
	"path/filepath"

	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"

	_ "github.com/joho/godotenv/autoload"
)

func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*") // Allow all origins
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Handle preflight request
		if r.Method == "OPTIONS" {
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	PORT := os.Getenv("PORT")

	initDB()
	defer db.Close()

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

	r.Use(enableCORS)
	handler := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"}, // Allow requests from this origin
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type"},
		AllowCredentials: true,
	}).Handler(r)
	// handler := cors.Default().Handler(r)
	handler = r

	log.Println("Server started at :" + PORT)
	if err := http.ListenAndServe(":"+PORT, handler); err != nil {
		log.Fatal(err)
	}
}
