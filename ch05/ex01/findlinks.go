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
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}

//!-main

//!+visit
// visit appends to links each link found in n and returns the result.
func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	first := n.FirstChild
	if first != nil {
		links = visit(links, first)
		n.RemoveChild(first)
		n.Type = html.ErrorNode
		links = visit(links, n)
	}

	return links
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
