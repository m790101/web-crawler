package main

import "net/url"

func normalizeUrl(urlString string) (string, error) {
	parsedURL, err := url.Parse(urlString)
	if err != nil {
		return "there is a err", err
	}
	hostName := parsedURL.Hostname()

	return hostName, nil
}
