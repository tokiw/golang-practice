package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"io"

	"bytes"

	"strings"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(fetch(os.Args[1]))
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	for _, text := range visit(nil, doc) {
		fmt.Println(text)
	}
}

//!-main

//!+visit
// visit appends to links each link found in n and returns the result.
func visit(text []string, n *html.Node) []string {
	if n.Type == html.TextNode && !isHiddenText(n) {
		for _, t := range strings.Split(n.Data, "\n") {
			if len(t) != 0 {
				text = append(text, n.Data)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		text = visit(text, c)
	}

	return text
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

func isHiddenText(n *html.Node) bool {
	parentData := n.Parent.Data
	if parentData == "script" || parentData == "style" {
		return true
	}
	return false
}
