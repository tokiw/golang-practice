package main

import (
	"fmt"
	"io"
	"os"

	"golang.org/x/net/html"
)

type Reader struct {
	s        string
	i        int64
	prevRune int
}

func (r *Reader) Read(p []byte) (n int, err error) {
	if r.i >= int64(len(r.s)) {
		return 0, io.EOF
	}
	r.prevRune = -1
	n = copy(p, r.s[r.i:])
	r.i += int64(n)
	return
}

func NewReader(s string) *Reader {
	return &Reader{s, 0, -1}
}

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
		</html>`
	reader := NewReader(src)
	doc, err := html.Parse(reader)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}

func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}
