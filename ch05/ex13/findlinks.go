package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
)

//!+breadthFirst
// breadthFirst calls f for each item in the worklist.
// Any items returned by f are added to the worklist.
// f is called at most once for each item.
func breadthFirst(f func(item string, hosts []string) []string, worklist []string) {
	hosts := make([]string, len(worklist))
	copy(hosts, worklist)
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item, hosts)...)
			}
		}
	}
}

//!-breadthFirst

//!+crawl
func crawl(u string, hosts []string) []string {
	fmt.Println(u)
	list, err := Extract(u)
	if err != nil {
		log.Print(err)
	}

	for _, link := range list {
		lurl, _ := url.Parse(link)
		for _, host := range hosts {
			hurl, _ := url.Parse(host)
			if hurl.Host == lurl.Host {
				makefile(lurl)
			}
		}
	}

	return list
}

func makefile(u *url.URL) {
	dir, filename, isDir := parseDir(u)
	if isDir {
		return
	}
	saveBasePath := "./links/" + u.Host + dir

	err := os.MkdirAll(saveBasePath, 0777)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	resp, err := http.Get(u.String())
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	if resp.StatusCode != 200 {
		fmt.Fprintln(os.Stdout, "redirect")
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	file, err := os.OpenFile(saveBasePath+filename, os.O_CREATE|os.O_WRONLY, 0666)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	file.Write(body)
	file.Close()
}

func parseDir(u *url.URL) (dir, filename string, isDir bool) {
	dir, filename = path.Split(u.Path)
	isDir = len(filename) == 0
	return
}

//!-crawl

//!+main
func main() {
	// Crawl the web breadth-first,
	// starting from the command-line arguments.
	breadthFirst(crawl, os.Args[1:])
}
