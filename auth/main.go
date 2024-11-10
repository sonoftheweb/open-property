package main

import (
	"log"
	"net/http"
	"open-property.com/auth/database"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	db, err := database.ConnectDB("host=auth_db user=authdb password=password dbname=authdb port=5432 sslmode=disable")
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	_ = db

	// Run migrations
	database.RunMigrations("host=auth_db user=authdb password=password dbname=authdb port=5432 sslmode=disable")

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})
	http.ListenAndServe(":8081", r)
}
