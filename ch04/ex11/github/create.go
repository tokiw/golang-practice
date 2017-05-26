package github

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

// CreateIssue the GitHub issue.
func CreateIssue(body io.Reader, accessToken string) (*Issue, error) {
	req, err := http.NewRequest(
		"POST",
		IssuesURL+"/repos/"+Repo+"/issues?",
		body,
	)
	if err != nil {
		return nil, err
	}

	// Content-Type 設定
	req.Header.Set("Authorization", "token "+accessToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed file open")
		return nil, err
	}

	fmt.Println(resp.StatusCode)

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
