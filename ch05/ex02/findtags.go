package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"io"

	"bytes"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(fetch(os.Args[1]))
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	for name, count := range visit(nil, doc) {
		fmt.Printf("%s: %d\n", name, count)
	}
}

//!-main

//!+visit
// visit appends to links each link found in n and returns the result.
func visit(tags map[string]int, n *html.Node) map[string]int {
	if tags == nil {
		tags = make(map[string]int)
	}
	if n.Type == html.ElementNode {
		tags[n.Data]++
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		tags = visit(tags, c)
	}

	return tags
}

func fetch(url string) io.Reader {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		os.Exit(1)
	}
	b, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
		os.Exit(1)
	}
	return bytes.NewReader(b)
}
