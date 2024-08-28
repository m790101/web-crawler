package main

import (
	"io"
	"net/http"
	"time"
)

func getHTML(url string) (string, error) {
	// make http call
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return "", err
	}

	client := http.Client{
		Timeout: time.Second * 10,
	}

	res, err := client.Do(req)

	if err != nil {
		return "", err
	}

	if res.Header.Get("Content-Type") != "text/html" {
		return "", nil
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		return "", nil
	}

	htmlBody, err := io.ReadAll(res.Body)

	if err != nil {
		return "", err
	}

	htmlString := string(htmlBody)

	return htmlString, nil
}
