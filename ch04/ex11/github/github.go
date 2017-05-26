package github

import "time"

const IssuesURL = "https://api.github.com"
const Repo = "tokiw/golang-practice"

type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string `json:"title"`
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string    `json:"body"`
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}
