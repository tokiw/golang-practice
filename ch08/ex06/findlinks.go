package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"gopl.io/ch5/links"
)

type linkInfo struct {
	url   string
	depth int
}

func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

//!+
func main() {
	depth := flag.Int("depth", 3, "max depthg")
	flag.Parse()
	worklist := make(chan []linkInfo)  // lists of URLs, may have duplicates
	unseenLinks := make(chan linkInfo) // de-duplicated URLs

	// Add command-line arguments to worklist.
	urls := os.Args[1:]

	go func() { worklist <- createlinkInfoList(urls, 0) }()

	// Create 20 crawler goroutines to fetch each unseen link.
	for i := 0; i < 20; i++ {
		go func() {
			for link := range unseenLinks {
				foundLinks := crawl(link.url)
				go func() { worklist <- createlinkInfoList(foundLinks, link.depth+1) }()
			}
		}()
	}

	// The main goroutine de-duplicates worklist items
	// and sends the unseen ones to the crawlers.
	seen := make(map[string]bool)
	for list := range worklist {
		for _, link := range list {
			if !seen[link.url] && (link.depth <= *depth) {
				seen[link.url] = true
				unseenLinks <- link
			}
		}
	}
}

func createlinkInfoList(urls []string, depth int) []linkInfo {
	var linkInfoList []linkInfo

	for _, url := range urls {
		linkInfoList = append(linkInfoList, linkInfo{url, depth})
	}

	return linkInfoList
}
