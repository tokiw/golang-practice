package main

import (
	"fmt"
	"os"

	"strings"

	"golang.org/x/net/html"
)

func main() {
	src := `<!DOCTYPE html>
		<html>
		<head>
		<title>title</title>
		</head>
		<body>
			<a href="hoge1.html">
			<a href="hoge2.html">
		</body>
		</html>
`
	doc, err := html.Parse(strings.NewReader(src))
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	for _, elem := range ElementsByTagName(doc, os.Args[1:]...) {
		fmt.Printf("<%s>\n", elem.Data)
	}
}

func ElementsByTagName(doc *html.Node, name ...string) []*html.Node {
	return visit(nil, doc, name)
}

//!+visit
// visit appends to links each link found in n and returns the result.
func visit(elements []*html.Node, n *html.Node, names []string) []*html.Node {
	if n.Type == html.ElementNode {
		if isTagElem(n, names) {
			elements = append(elements, n)
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		elements = visit(elements, c, names)
	}

	return elements
}

func isTagElem(n *html.Node, tags []string) bool {
	for _, tag := range tags {
		if n.Data == tag {
			return true
		}
	}
	return false
}
