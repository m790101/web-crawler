package main

import (
	"testing"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestNormalizeUrl(t *testing.T) {
	urlString, err := normalizeUrl("https://www.google.com/path")
	if err != nil {
		t.Fatalf("Expected no error, but got %s", err)
	}
	if urlString != "www.google.com" {
		t.Fatalf("Expected www.google.com, but got %s", urlString)
	}
}
