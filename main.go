package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		log.Println("Not enough arguments provided")
		os.Exit(1)
	}

	page := args[1]
	html, err := getHTML(page)
	if err != nil {
		log.Println("Error parsing html: ", err)
	}

	getLinks(page, html)
	miracles := getMiracles(html)

	fmt.Printf("Miracles: %q\n", miracles)
}
