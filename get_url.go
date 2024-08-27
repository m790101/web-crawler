package main

import (
	"strings"

	"golang.org/x/net/html"
)

func getURLsFromHTML(htmlBody, rawBaseURL string) ([]string, error) {
	doc, err := html.Parse(strings.NewReader(htmlBody))

	if err != nil {
		return nil, err
	}

	urls := getURLsFromNode(doc, []string{}, &rawBaseURL)

	return urls, nil
}

func getURLsFromNode(node *html.Node, urls []string, rawBaseURL *string) []string {

	if node.Type == html.ElementNode && node.Data == "a" {
		for _, a := range node.Attr {
			if a.Key == "href" {
				rawBaseURL := *rawBaseURL
				if !strings.HasPrefix(a.Val, "https://") {
					a.Val = rawBaseURL + a.Val
				}
				urls = append(urls, a.Val)
			}
		}
	}

	if node.FirstChild != nil {
		urls = getURLsFromNode(node.FirstChild, urls, rawBaseURL)
	}
	if node.NextSibling != nil {
		urls = getURLsFromNode(node.NextSibling, urls, rawBaseURL)
	}

	return urls
}
