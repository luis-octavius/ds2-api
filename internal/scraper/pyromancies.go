package main

import (
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func getPyromancies(html string) []Pyromancy {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		log.Println("getPyromancies - error: ", err)
		return []Pyromancy{}
	}

	pyromancies := make([]Pyromancy, 0)

	doc.Find("table.wiki_table").Each(func(tableIdx int, table *goquery.Selection) {
		table.Find("tbody tr").Each(func(rowIdx int, row *goquery.Selection) {
			debugRow(row)
			pyromancy := extractPyromancyFromRow(row)

			if pyromancy.Name != "" {
				pyromancies = append(pyromancies, pyromancy)
			}
		})
	})

	return pyromancies
}

func extractPyromancyFromRow(row *goquery.Selection) Pyromancy {
	pyro := Pyromancy{}

	cells := row.Find("td")

	if cells.Length() < 7 {
		return Pyromancy{}
	}

	pyro.Name = cleanText(cells.Eq(0))
	pyro.Uses = cleanText(cells.Eq(1))
	pyro.Attunement = cleanText(cells.Eq(2))
	pyro.Description = cleanText(cells.Eq(3))
	pyro.Acquired = cleanText(cells.Eq(4))
	pyro.Cost = cleanText(cells.Eq(5))
	pyro.Type = cleanText(cells.Eq(6))

	return pyro

}
