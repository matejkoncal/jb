package main

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"

	"github.com/andygrunwald/go-jira"
)

func main() {
	login := getLogin()
	issueNumber := os.Args[1]

	tp := jira.BasicAuthTransport{
		Username: login.Username,
		Password: login.Token,
	}

	client, err := jira.NewClient(tp.Client(), login.URL)

	issue, _, err := client.Issue.Get(issueNumber, nil)

	if err != nil {
		panic(err)
	}

	branch := sanitizeBranchName(issue.Key + "-" + issue.Fields.Summary)

	checkoutToBranch(branch)
}

func sanitizeBranchName(branchName string) string {
	sanitized := strings.ReplaceAll(strings.ToLower(branchName), " ", "-")
	re := regexp.MustCompile("[^a-zA-Z0-9-]")

	withDashes := re.ReplaceAllString(sanitized, "-")
	return strings.Trim(regexp.MustCompile(`-+`).ReplaceAllString(withDashes, "-"), "-")
}

func checkoutToBranch(branch string) {
	cmd := exec.Command("git", "checkout", branch)
	output, err := cmd.CombinedOutput()

	if err != nil {
		cmd := exec.Command("git", "checkout", "-b", branch)
		output, err := cmd.CombinedOutput()

		if err != nil {
			panic(err)
		}

		fmt.Print(string(output))
	} else {
		fmt.Print(string(output))
	}
}

type Login struct {
	Username string
	Token    string
	URL      string
}

func getLogin() Login {
	return Login{
		Username: os.Getenv("JIRA_USERNAME"),
		Token:    os.Getenv("JIRA_TOKEN"),
		URL:      os.Getenv("JIRA_URL"),
	}
}
