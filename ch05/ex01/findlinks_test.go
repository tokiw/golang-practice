package main

import (
	"fmt"
	"os"
	"testing"

	"golang.org/x/net/html"
)

func TestFindLinks(t *testing.T) {
	doc, err := html.Parse(fetch("https://golang.org"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	links := visit(nil, doc)
	expected := []string{
		"/",
		"/",
		"#",
		"/doc/",
		"/pkg/",
		"/project/",
		"/help/",
		"/blog/",
		"http://play.golang.org/",
		"#",
		"#",
		"//tour.golang.org/",
		"https://golang.org/dl/",
		"//blog.golang.org/",
		"https://developers.google.com/site-policies#restrictions",
		"/LICENSE",
		"/doc/tos.html",
		"http://www.google.com/intl/en/policies/privacy/",
	}

	assertEqual(expected, links, t)
}

func assertEqual(expected []string, actual []string, t *testing.T) {
	if len(expected) != len(actual) {
		t.Error("not same size")
		return
	}
	for i := range expected {
		if expected[i] != actual[i] {
			t.Errorf("not same value. expected is %v, but actual is %v", expected[i], actual[i])
			return
		}
	}
}
