package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	words, images, err := CountWordsAndImages(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
	fmt.Printf("words: %d images: %d", words, images)
}

// CountWordsAndImages count words and images
func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		return
	}
	words, images = countWordsAndImages(doc)
	return
}

func countWordsAndImages(n *html.Node) (words, images int) {
	return visit(0, 0, n)
}

func visit(w int, i int, n *html.Node) (int, int) {
	if n.Type == html.TextNode && !isHiddenText(n) {
		for _, t := range strings.Split(n.Data, "\n") {
			if len(t) != 0 {
				w += countWords(n.Data)
			}
		}
	} else if n.Type == html.ElementNode && n.Data == "img" {
		for _, a := range n.Attr {
			if a.Key == "src" {
				i++
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		w, i = visit(w, i, c)
	}
	return w, i
}

func isHiddenText(n *html.Node) bool {
	parentData := n.Parent.Data
	if parentData == "script" || parentData == "style" {
		return true
	}
	return false
}

func countWords(s string) int {
	input := bufio.NewScanner(strings.NewReader(s))
	input.Split(bufio.ScanWords)

	var count int
	for input.Scan() {
		input.Text()
		count++
	}
	return count
}
