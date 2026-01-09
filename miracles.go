package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Miracle struct {
	Name        string
	Slots       string
	Faith       string
	Attunement  string
	Description string
	Acquired    string
	Type        string
}

func getLinks(rawURL string, html string) []string {
	links := []string{}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		log.Println("getLinks - error: ", err)
		return links
	}

	doc.Find("a.wiki_link").Each(func(_ int, s *goquery.Selection) {
		text := s.Text()
		link, ok := s.Attr("href")
		if ok {
			constructedURL := rawURL + link
			links = append(links, constructedURL)
		}
		fmt.Println(text)
	})

	return links
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
			debugRow(row)

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
	miracle.Slots = cleanText(cells.Eq(1))
	miracle.Faith = cleanText(cells.Eq(2))
	miracle.Attunement = cleanText(cells.Eq(3))
	miracle.Description = cleanText(cells.Eq(4))
	miracle.Acquired = cleanText(cells.Eq(5))
	miracle.Type = cleanText(cells.Eq(6))

	return miracle
}

func cleanText(sel *goquery.Selection) string {
	text := sel.Text()

	text = strings.TrimSpace(text)
	text = strings.Join(strings.Fields(text), " ")

	text = strings.Trim(text, "-\u00a0")

	return text
}

func cleanDescription(sel *goquery.Selection) string {
	text := sel.Text()
	text = strings.TrimSpace(text)

	text = strings.ReplaceAll(text, "\n", " ")
	text = strings.Join(strings.Fields(text), " ")

	return text
}
