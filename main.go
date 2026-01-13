package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

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

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

	queries := database.New(db)
	apiCfg.db = queries

	mux := http.NewServeMux()

	server := http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	// magic endpoints
	mux.HandleFunc("GET /api/magic/miracles", apiCfg.handlerGetMiracles)
	mux.HandleFunc("GET /api/magic/sorceries", apiCfg.handlerGetSorceries)
	mux.HandleFunc("GET /api/magic/pyromancies", apiCfg.handlerGetPyromancies)
	mux.HandleFunc("GET /api/magic/hexes", apiCfg.handlerGetHexes)

	fmt.Printf("Server running on port %s\n", port)
	err = server.ListenAndServe()
	if err != nil {
		log.Fatalf("Error listening server on port %s: %v", port, err)
	}

}
