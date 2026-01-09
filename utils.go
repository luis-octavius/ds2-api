package main

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func cleanText(sel *goquery.Selection) string {
	text := sel.Text()

	text = strings.TrimSpace(text)
	text = strings.Join(strings.Fields(text), " ")

	text = strings.Trim(text, "-\u00a0")

	return text
}
