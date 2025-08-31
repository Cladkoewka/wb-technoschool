package main

import (
	"fmt"
	"os"

	"github.com/Cladkoewka/wb-technoschool/level2/task16/internal/crawling"
)


func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: wget <url> <depth>")
		return
	}
	rawURL := os.Args[1]
	depth := 1
	fmt.Sscanf(os.Args[2], "%d", &depth)

	crawler, err := crawling.NewCrawler(rawURL, depth, "crawled")
	if err != nil {
		fmt.Fprintf(os.Stderr, "init error: %v\n", err)
		os.Exit(1)
	}
	crawler.Start()
}