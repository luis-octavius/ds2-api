package main

import (
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Hex struct {
	Name         string
	Uses         string
	Attunement   string
	Intelligence string
	Faith        string
	Description  string
	Acquired     string
	Cost         string
	Type         string
}

func getHexes(html string) []Hex {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		log.Println("getHexes - error: ", err)
		return []Hex{}
	}

	hexes := make([]Hex, 0)

	doc.Find("table.wiki_table").Each(func(tableIdx int, table *goquery.Selection) {
		table.Find("tbody tr").Each(func(rowIdx int, row *goquery.Selection) {
			if row.Find("th").Length() > 0 {
				return
			}

			hex := extractHexFromRow(row)

			if hex.Name != "" {
				hexes = append(hexes, hex)
			}
		})
	})

	return hexes
}

func extractHexFromRow(row *goquery.Selection) Hex {
	hex := Hex{}

	cells := row.Find("td")

	hex.Name = cleanText(cells.Eq(0))
	hex.Uses = cleanText(cells.Eq(1))
	hex.Attunement = cleanText(cells.Eq(2))
	hex.Intelligence = cleanText(cells.Eq(3))
	hex.Faith = cleanText(cells.Eq(4))
	hex.Description = cleanText(cells.Eq(5))
	hex.Acquired = cleanText(cells.Eq(6))
	hex.Cost = cleanText(cells.Eq(7))
	hex.Type = cleanText(cells.Eq(8))

	return hex
}
