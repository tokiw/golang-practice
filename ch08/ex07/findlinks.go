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
	worklist := make(chan []string)
	hosts := os.Args[1:]
	var n int
	n++
	go func() {
		worklist <- os.Args[1:]
	}()

	seen := make(map[string]bool)
	for ; n > 0; n-- {
		list := <-worklist
		for _, item := range list {
			if !seen[item] {
				seen[item] = true
				n++
				go func(item string) {
					worklist <- crawl(item, hosts)
				}(item)
			}
		}
	}
}
