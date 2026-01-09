package main

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func debugRow(row *goquery.Selection) {
	fmt.Println("\n=== DEBUG ROW ===")

	row.Find("td").Each(func(idx int, cell *goquery.Selection) {
		html, _ := cell.Html()
		text := cell.Text()

		fmt.Printf("Cell %d:\n", idx)
		fmt.Printf("   HTML: %s\n", strings.ReplaceAll(html[:min(100, len(html))], "\n", "\\n"))
		fmt.Printf("  TEXT: %q\n", text)
		fmt.Printf("   Lines: %v\n\n", strings.Split(text, "\n"))
	})
}
