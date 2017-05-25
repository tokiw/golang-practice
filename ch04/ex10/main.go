package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/tokiw/golang-practice/ch04/ex10/github"
)

//!+
func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	now := time.Now()
	var lessThanMonth []*github.Issue
	var lessThanYear []*github.Issue
	var moreThanYear []*github.Issue

	for _, item := range result.Items {
		if item.CreatedAt.After(now.Add(-30 * 24 * time.Hour)) {
			lessThanMonth = append(lessThanMonth, item)
		} else if item.CreatedAt.After(now.Add(-365 * 24 * time.Hour)) {
			lessThanYear = append(lessThanYear, item)
		} else {
			moreThanYear = append(moreThanYear, item)
		}
	}

	fmt.Printf("%d issues in less than a month:\n", len(lessThanMonth))
	for _, item := range lessThanMonth {
		fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
	}
	fmt.Printf("%d issues in less than a year old:\n", len(lessThanYear))
	for _, item := range lessThanYear {
		fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
	}
	fmt.Printf("%d issues in more than a year old:\n", len(moreThanYear))
	for _, item := range moreThanYear {
		fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
	}
}
