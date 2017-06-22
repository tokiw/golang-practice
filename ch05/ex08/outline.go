// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 133.

// Outline prints the outline of an HTML document tree.
package main

import (
	"fmt"
	"net/http"

	"os"

	"golang.org/x/net/html"
)

func main() {
	resp, err := http.Get("https://golang.org")
	if err != nil {
		os.Exit(1)
	}

	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		os.Exit(1)
	}
	if len(os.Args) < 2 {
		os.Exit(1)
	}

	node := ElementByID(doc, os.Args[1])
	fmt.Println(node.Data)
}

func outline(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return err
	}

	//!+call
	forEachNode(doc, startElement, endElement)
	//!-call

	return nil
}

//!+forEachNode
// forEachNode calls the functions pre(x) and post(x) for each node
// x in the tree rooted at n. Both functions are optional.
// pre is called before the children are visited (preorder) and
// post is called after (postorder).
func forEachNode(n *html.Node, pre, post func(n *html.Node) bool) {
	if pre != nil && !pre(n) {
		return
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil && post(n) {
		return
	}
}

//!-forEachNode

func ElementByID(doc *html.Node, id string) *html.Node {
	var node *html.Node
	forEachNode(doc, func(n *html.Node) bool {
		if n.Type == html.ElementNode {
			for _, attr := range n.Attr {
				if attr.Key == "id" && attr.Val == id {
					node = n
					return false
				}
			}
		}
		return startElement(n)
	}, endElement)
	return node
}

//!+startend
var depth int

func startElement(n *html.Node) bool {
	if n.Type == html.ElementNode {
		// fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
		depth++
	}
	return true
}

func endElement(n *html.Node) bool {
	if n.Type == html.ElementNode {
		depth--
		// fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
		if n.Data == "html" {
			return false
		}
	}
	return true
}

//!-startend
