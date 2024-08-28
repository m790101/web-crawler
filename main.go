package main

import (
	"fmt"
	"os"
)

func main() {

	argsWithoutProg := os.Args
	if len(argsWithoutProg) < 2 {
		fmt.Println("no website provided")
		os.Exit(1)
	}
	if len(argsWithoutProg) > 2 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}

	rawBaseURL := argsWithoutProg[1]

	htmlString, _ := getHTML(rawBaseURL)
	fmt.Println(htmlString)

}
