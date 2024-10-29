package main

import (
	"chapter_4/4.11/github"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

// Your Github token
var token = "GITHUB TOKEN"

func search(query []string) {
	result, err := github.SearchIssues(query)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		fmt.Printf("#%-5d %20s %.55q\n", item.Number, item.User.Login, item.Title)
	}
}

func read(owner, repo, number string)  {
	item, err := github.GetIssue(owner, repo, number)
	if err != nil {
		log.Fatal(err)
	}
	body := item.Body
	if body == "" {
		body = "<NO BODY>\n"
	}
	fmt.Printf("repo: %s/%s\nnumber: %s\nuser: %s\nstate: %s\ntitle: %s\n\n%s",
		owner, repo, number, item.User.Login, item.State, item.Title, body)
}

func create(owner, repo string) {
	title, body, err := getInputFromEditor()
	if err != nil {
		log.Fatal(err)
	}

	issue, err := github.CreateIssue(owner, repo, title, body, token)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Created issue #%d: %s\n", issue.Number, issue.HTMLURL)
}

func update(owner, repo, number string) {
	title, body, err := getInputFromEditor()
	if err != nil {
		log.Fatal(err)
	}

	issue, err := github.UpdateIssue(owner, repo, number, title, body, token)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Updated issue #%d: %s\n", issue.Number, issue.HTMLURL)
}

func close(owner, repo, number string) {
	err := github.CloseIssue(owner, repo, number, token)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Closed issue #%s\n", number)
}

func getInputFromEditor() (string, string, error) {
	editor := `C:\Program Files\Vim\vim91\gvim.exe`

	tmpfile, err := os.CreateTemp("", "issue*.txt")
	if err != nil {
		return "", "", err
	}
	defer os.Remove(tmpfile.Name())

	cmd := exec.Command(editor, tmpfile.Name())
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return "", "", err
	}

	content, err := os.ReadFile(tmpfile.Name())
	if err != nil {
		return "", "", err
	}

	parts := strings.SplitN(string(content), "\n", 2)
	title := strings.TrimSpace(parts[0])
	body := ""
	if len(parts) > 1 {
		body = strings.TrimSpace(parts[1])
	}
	return title, body, nil
}



var usage string = `usage: 
search QUERY
[create|read|update|delete] OWNER REPO ISSUE_NUMBER
`

func usageDie()  {
	fmt.Fprintln(os.Stderr, usage)
	os.Exit(1)
}

func main() {
	if len(os.Args) < 2 {
		usageDie()
	}
	action := os.Args[1]
	args := os.Args[2:]

	if action == "search" {
		if len(args) < 1 {
			usageDie()
		}
		search(args)
		os.Exit(0)
	}
	if len(args) < 2 {
		usageDie()
	}

	owner, repo := args[0], args[1]
	if action == "create" {
		create(owner, repo)
	} else if action == "read" && len(args) == 3{
		number := args[2]
		read(owner, repo, number)
	} else if action == "update" && len(args) == 3{
		number := args[2]
		update(owner, repo, number)
	} else if action == "close" && len(args) == 3 {
		number := args[2]
		close(owner, repo, number)
	} else {
		usageDie()
	}
}