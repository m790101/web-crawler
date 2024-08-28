package main

import (
	"fmt"
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

func getURLsFromHTML(htmlBody, rawBaseURL string) ([]string, error) {

	baseURL, err := url.Parse(rawBaseURL)
	if err != nil {

		return nil, fmt.Errorf("couldn't parse base URL")
	}

	doc, err := html.Parse(strings.NewReader(htmlBody))

	if err != nil {
		return nil, fmt.Errorf("couldn't parse HTML")
	}

	urls, _ := getURLsFromNode(doc, []string{}, baseURL)

	return urls, nil
}

func getURLsFromNode(node *html.Node, urls []string, baseURL *url.URL) ([]string, error) {

	if node.Type == html.ElementNode && node.Data == "a" {
		for _, a := range node.Attr {
			if a.Key == "href" {
				href, err := url.Parse(a.Val)
				if err != nil {
					fmt.Printf("couldn't parse href '%v': %v\n", a.Val, err)
					continue
				}
				resolvedURL := baseURL.ResolveReference(href)
				urls = append(urls, resolvedURL.String())

			}
		}
	}

	if node.FirstChild != nil {
		urls, _ = getURLsFromNode(node.FirstChild, urls, baseURL)
	}
	if node.NextSibling != nil {
		urls, _ = getURLsFromNode(node.NextSibling, urls, baseURL)
	}
	if len(urls) == 0 {
		return nil, nil
	}
	return urls, nil
}
