package main

import (
	"database/sql"
	"log"
	"net/url"

	_ "github.com/lib/pq"
	"github.com/luis-octavius/ds2-scraper/internal/database"
)

type state struct {
	db      *database.Queries
	pages   map[string]string
	baseURL *url.URL
}

func main() {
	// args validation
	// args := os.Args
	// if len(args) < 2 {
	// 	log.Println("Not enough arguments provided")
	// 	os.Exit(1)
	// }
	//
	dbUrl := "postgres://postgres:@localhost:5432/ds2?sslmode=disable"
	// baseUrl := args[1]
	// parsedURL, err := url.Parse(baseUrl)
	// if err != nil {
	// 	log.Fatal("Error parsing URL: ", err)
	// }

	db, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatal("Error opening the database: ", err)
	}

	queries := database.New(db)

	// state of the scraper
	state := state{
		db: queries,
	}

	htmlMiracles, err := getHTML("https://darksouls2.wiki.fextralife.com/Miracles")
	if err != nil {
		log.Fatal("Error extracting html from miracles URL: ", err)
	}

	miracles := getMiracles(htmlMiracles)
	state.SaveMiracles(miracles)

	htmlPyromancies, err := getHTML("https://darksouls2.wiki.fextralife.com/Pyromancies")
	if err != nil {
		log.Fatal("Error extracting html from pyromancies URL: ", err)
	}

	pyromancies := getPyromancies(htmlPyromancies)
	state.SavePyromancies(pyromancies)

	htmlHexes, err := getHTML("https://darksouls2.wiki.fextralife.com/Hexes")
	if err != nil {
		log.Fatal("Error extracting html from hexes URL: ", err)
	}

	hexes := getHexes(htmlHexes)
	state.SaveHexes(hexes)

	htmlSorceries, err := getHTML("https://darksouls2.wiki.fextralife.com/Sorceries")
	if err != nil {
		log.Fatal("Error extracting html from sorceries URL: ", err)
	}

	sorceries := getSorceries(htmlSorceries)
	state.SaveSorceries(sorceries)
}
