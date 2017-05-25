package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type poster struct {
	Title  string
	Poster string
}

func main() {
	result, err := searchPoster(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	downloadPoster(*result)
}

func searchPoster(keywords []string) (*poster, error) {
	q := url.QueryEscape(strings.Join(keywords, "+"))
	resp, err := http.Get("http://www.omdbapi.com/?t=" + q)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	var result poster
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

func downloadPoster(p poster) {
	//
}
