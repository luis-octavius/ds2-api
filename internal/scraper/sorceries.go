package main

import (
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func getSorceries(html string) []Sorcery {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		log.Println("getSorceries - error: ", err)
		return []Sorcery{}
	}

	sorceries := []Sorcery{}

	doc.Find("table.wiki_table").Each(func(tableIdx int, table *goquery.Selection) {
		doc.Find("tbody tr").Each(func(rowIdx int, row *goquery.Selection) {
			sorcery := extractSorceryFromRow(row)

			if sorcery.Name != "" {
				sorceries = append(sorceries, sorcery)
			}
		})
	})

	return sorceries
}

func extractSorceryFromRow(row *goquery.Selection) Sorcery {
	sorcery := Sorcery{}

	cells := row.Find("td")

	if cells.Length() < 8 {
		return sorcery
	}

	sorcery.Name = cleanText(cells.Eq(0))
	sorcery.Uses = cleanText(cells.Eq(1))
	sorcery.Spellslot = cleanText(cells.Eq(2))
	sorcery.Intelligence = cleanText(cells.Eq(3))
	sorcery.Description = cleanText(cells.Eq(4))
	sorcery.Acquired = cleanText(cells.Eq(5))
	sorcery.Cost = cleanText(cells.Eq(6))
	sorcery.Type = cleanText(cells.Eq(7))

	return sorcery
}
