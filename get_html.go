package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func getHTML(rawURL string) (string, error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", rawURL, nil)
	if err != nil {
		return "", fmt.Errorf("getHTML - error creating the request with provided URL %v: %v", rawURL, err)
	}

	req.Header.Set("User-Agent", "OctaviusCrawler/1.0")

	res, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("getHTML - error doing the request: %v", err)
	}

	defer res.Body.Close()

	if res.StatusCode >= 400 {
		return "", fmt.Errorf("getHTML - status code not successful: %v", res.StatusCode)
	}

	contentType := res.Header.Get("Content-Type")
	if !strings.Contains(contentType, "text/html") {
		return "", fmt.Errorf("getHTML - content type is not text/html: %v", contentType)
	}

	body, err := io.ReadAll(res.Body)

	return string(body), nil
}
