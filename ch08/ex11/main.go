package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
)

type response struct {
	filename string
	n        int64
	err      error
}

//!+
// Fetch downloads the URL and returns the
// name and length of the local file.
func fetch(url string, cancel <-chan struct{}) (filename string, n int64, err error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", 0, err
	}
	req.Cancel = cancel
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()

	local := path.Base(resp.Request.URL.Path)
	if local == "/" {
		local = "index.html"
	}
	f, err := os.Create(local)
	if err != nil {
		return "", 0, err
	}
	n, err = io.Copy(f, resp.Body)
	// Close file, but prefer error from Copy, if any.
	if closeErr := f.Close(); err == nil {
		err = closeErr
	}
	return local, n, err
}

//!-

func main() {
	cancel := make(chan struct{})
	resp := make(chan response)
	for _, url := range os.Args[1:] {
		go func(url string) {
			local, n, err := fetch(url, cancel)
			if err != nil {
				fmt.Fprintf(os.Stderr, "fetch %s: %v\n", url, err)
				return
			}
			close(cancel)
			resp <- response{local, n, err}
		}(url)
	}
	res := <-resp
	fmt.Printf(" %s (%d bytes).\n", res.filename, res.n)
}
