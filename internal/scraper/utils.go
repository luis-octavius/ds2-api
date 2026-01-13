package main

import (
	"context"
	"net/url"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

var ctx = context.Background()

func cleanText(sel *goquery.Selection) string {
	text := sel.Text()

	text = strings.TrimSpace(text)
	text = strings.Join(strings.Fields(text), " ")

	text = strings.Trim(text, "-\u00a0")

	return text
}

func constructURL(baseURL *url.URL, path string) (*url.URL, error) {
	constructedURL := baseURL.String() + "/" + path

	parsedURL, err := url.Parse(constructedURL)
	if err != nil {
		return nil, err
	}
	return parsedURL, nil
}

func strToInt(num string) int32 {
	convertedNum, _ := strconv.Atoi(num)
	return (int32)(convertedNum)
}
