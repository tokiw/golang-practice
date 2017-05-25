package github

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

// CreateIssue the GitHub issue.
func CreateIssue(body io.Reader) (*Issue, error) {
	resp, err := http.Post(IssuesURL+"/repos/"+Repo+"/issues", "application/json", body)

	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed file open")
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, err
	}

	var result Issue
	if err = json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil
}
