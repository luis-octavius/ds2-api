package main

import (
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Pyromancy struct {
	Name        string
	Uses        string
	Attunement  string
	Description string
	Acquired    string
	Cost        string
	Type        string
}

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

type Miracle struct {
	Name        string
	Uses        string
	Faith       string
	Attunement  string
	Description string
	Acquired    string
	Type        string
}

type Sorcery struct {
	Name         string
	Uses         string
	Spellslot    string
	Intelligence string
	Description  string
	Acquired     string
	Cost         string
	Type         string
}

func getMiracles(html string) []Miracle {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		log.Println("getMiracles - error: ", err)
	}

	miracles := []Miracle{}
	doc.Find("table.wiki_table").Each(func(tableIdx int, table *goquery.Selection) {
		table.Find("tbody tr").Each(func(rowIdx int, row *goquery.Selection) {
			// skip header rows that have th instead of td
			if row.Find("th").Length() > 0 {
				return
			}

			miracle := extractMiracleFromRow(row)

			if miracle.Name != "" {
				miracles = append(miracles, miracle)
			}
		})
	})

	return miracles
}

func extractMiracleFromRow(row *goquery.Selection) Miracle {
	miracle := Miracle{}

	cells := row.Find("td")

	if cells.Length() < 7 {
		return miracle
	}

	miracle.Name = cleanText(cells.Eq(0))
	miracle.Uses = cleanText(cells.Eq(1))
	miracle.Attunement = cleanText(cells.Eq(2))
	miracle.Faith = cleanText(cells.Eq(3))
	miracle.Description = cleanText(cells.Eq(4))
	miracle.Acquired = cleanText(cells.Eq(5))
	miracle.Type = cleanText(cells.Eq(6))

	return miracle
}

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

func getPyromancies(html string) []Pyromancy {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		log.Println("getPyromancies - error: ", err)
		return []Pyromancy{}
	}

	pyromancies := make([]Pyromancy, 0)

	doc.Find("table.wiki_table").Each(func(tableIdx int, table *goquery.Selection) {
		table.Find("tbody tr").Each(func(rowIdx int, row *goquery.Selection) {
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
