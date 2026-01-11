package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/luis-octavius/ds2-scraper/internal/database"
)

type apiConfig struct {
	db *database.Queries
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	apiCfg := apiConfig{}

	dbUrl := os.Getenv("DB_URL")
	port := os.Getenv("PORT")

	db, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatalf("Error opening the database with the provided url %s: %v", dbUrl, err)
	}

	queries := database.New(db)
	apiCfg.db = queries

	mux := http.NewServeMux()

	server := http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	err = server.ListenAndServe()
	if err != nil {
		log.Fatalf("Error listening server on port %s: %v", port, err)
	}

}
