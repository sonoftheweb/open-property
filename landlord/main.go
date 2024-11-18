package main

import (
	"log"
	"net/http"
	"os"

	"open-property.com/landlord/database"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Get connection string from environment variable
	connectionString := os.Getenv("CONNECTION_STRING")
	if connectionString == "" {
		log.Fatalf("CONNECTION_STRING environment variable not set")
	}

	db, err := database.ConnectDB(connectionString)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	_ = db

	// Run migrations
	database.RunMigrations(connectionString)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"message": "Hello World!"}`))
	})

	// r.Post("/register", createUserHandler(db))

	http.ListenAndServe(":8082", r)
}
