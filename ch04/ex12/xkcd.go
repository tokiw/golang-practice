package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type Comic struct {
	Month      string
	Num        int `json:"num"`
	Link       string
	Year       string
	News       string
	Transcript string // in Markdown format
	SafeTitle  string `json:"safe_title"`
	Alt        string
	Image      string `json:"img"`
	Day        string
}

func main() {
	num, err := getLatestNum()

	if err != nil {
		log.Fatal(err)
	}

	if num < 0 {
		fmt.Errorf("Error")
	}
	files, _ := fetchAll(num)
	q := os.Args[1]

	var comics []Comic
	for _, file := range files {
		c := parse(file)
		if find(c, q) {
			comics = append(comics, c)
		}
	}
	for _, c := range comics {
		fmt.Printf("%v: %v(%v)", c.SafeTitle, c.Transcript, c.Link)
	}
}

func getLatestNum() (int, error) {
	result, err := http.Get("https://xkcd.com/info.0.json")

	if err != nil {
		return -1, err
	}

	if result.StatusCode != http.StatusOK {
		result.Body.Close()
		return -1, fmt.Errorf("fetch failed: %s", result.Status)
	}

	var latest Comic
	if err := json.NewDecoder(result.Body).Decode(&latest); err != nil {
		result.Body.Close()
		return -1, err
	}
	result.Body.Close()
	return latest.Num, nil
}

func fetchAll(num int) ([]string, error) {
	dir := "./data"
	idToReplace := "{ID}"
	url := "https://xkcd.com/" + idToReplace + "/info.0.json"

	for i := 1; i <= num; i++ {
		path := dir + "/" + strconv.Itoa(i) + ".json"
		_, err := os.Stat(path)
		if err == nil {
			continue
		}
		resp, err := http.Get(strings.Replace(url, idToReplace, strconv.Itoa(i), 1))
		if err != nil {
			return nil, err
		}

		if resp.StatusCode != http.StatusOK {
			resp.Body.Close()
			fmt.Println(" ERROR!")
			return nil, fmt.Errorf("fetch failed: %s", resp.Status)
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		err = ioutil.WriteFile(path, body, 0644)
		if err != nil {
			return nil, err
		}
		resp.Body.Close()
	}

	files, _ := filepath.Glob(dir + "/*")
	return files, nil
}

func find(c Comic, query string) bool {
	if strings.Contains(c.SafeTitle, query) {
		return true
	} else if strings.Contains(c.Transcript, query) {
		return true
	}
	return false
}

func parse(file string) Comic {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
	var c Comic
	if err := json.Unmarshal(data, &c); err != nil {
		log.Fatal(err)
	}
	return c
}
