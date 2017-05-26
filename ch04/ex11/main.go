package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"flag"

	"github.com/tokiw/golang-practice/ch04/ex11/github"
)

func main() {
	if len(os.Args) <= 1 {
		log.Fatal("Specify a command")
		return
	}
	command := os.Args[1]

	if len(command) <= 0 {
		log.Fatal("Specify a command")
		return
	}

	switch command {
	case "search":
		search()
	case "create":
		create()
	}
}

// search repo:golang/go is:open json decoder
func search() {
	result, err := github.SearchIssues()
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}
}

func create() {
	sub := os.Args[1]
	f := flag.NewFlagSet(os.Args[0]+" "+sub, flag.ExitOnError)
	title := f.String("title", "111", "Issue Title")
	body := f.String("body", "", "Issue Body")
	token := f.String("token", "", "Access Token")
	f.Parse(os.Args[2:])

	var issue github.Issue
	issue.Title = *title
	issue.Body = *body

	json, _ := json.Marshal(issue)
	result, err := github.CreateIssue(bytes.NewBuffer(json), *token)
	if err != nil {
		log.Fatal(err)
		return
	}
	if result == nil {
		log.Fatal("Error")
		return
	}
	fmt.Printf("Created issue: %v\n", result.Title)
}
