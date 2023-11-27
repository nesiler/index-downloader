// Author: Enes Diler | Github: @nesiler | Website: nesiler.com | mail: me@nesiler.com

package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run *.go <url>")
		os.Exit(1)
	}

	url := os.Args[1]

	scraper, err := newScraper(url)
	if err != nil {
		fmt.Println("Error creating scraper:", err)
		os.Exit(1)
	}

	scraper.Start()
}
