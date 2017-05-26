package github

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// CreateIssue the GitHub issue.
func CreateIssue(body io.Reader, accessToken string) error {
	req, err := http.NewRequest(
		"POST",
		IssuesURL+"/repos/"+Repo+"/issues?",
		body,
	)
	if err != nil {
		return err
	}

	// Content-Type 設定
	req.Header.Set("Authorization", "token "+accessToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed file open")
		return err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return err
	}
	fmt.Println("Create!!")
	return nil
}
