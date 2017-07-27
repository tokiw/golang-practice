package main

import (
	"fmt"
	"log"
	"os"
)

func crawl(url string, cancel <-chan struct{}) []string {
	fmt.Println(url)
	list, err := Extract(url, cancel)
	if err != nil {
		log.Print(err)
	}
	return list
}

//!+
func main() {
	worklist := make(chan []string)  // lists of URLs, may have duplicates
	unseenLinks := make(chan string) // de-duplicated URLs
	cancel := make(chan struct{})

	// Add command-line arguments to worklist.
	go func() { worklist <- os.Args[1:] }()
	go func() {
		os.Stdin.Read(make([]byte, 1))
		close(cancel)
		close(worklist) // ちょっといけてない
	}()

	// Create 20 crawler goroutines to fetch each unseen link.
	for i := 0; i < 20; i++ {
		go func() {
			for link := range unseenLinks {
				foundLinks := crawl(link, cancel)
				go func() { worklist <- foundLinks }()
			}
		}()
	}

	// The main goroutine de-duplicates worklist items
	// and sends the unseen ones to the crawlers.
	seen := make(map[string]bool)
	for list := range worklist {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				unseenLinks <- link
			}
		}
	}
}
