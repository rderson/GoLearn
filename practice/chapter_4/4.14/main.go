package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

type Issue struct {
	Title   string `json:"title"`
	Number  int    `json:"number"`
	User    *User  `json:"user"`
	HTMLURL string `json:"html_url"`
}

type User struct {
	Login   string `json:"login"`
	HTMLURL string `json:"html_url"`
}

type Milestone struct {
	Title  string `json:"title"`
	Number int    `json:"number"`
	HTMLURL string `json:"html_url"`
}

func fetchGitHubData(owner, repo string) ([]Issue, []Milestone, error) {
	var issues []Issue
	var milestones []Milestone

	issuesURL := fmt.Sprintf("https://api.github.com/repos/%s/%s/issues", owner, repo)
	resp, err := http.Get(issuesURL)
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&issues); err != nil {
		return nil, nil, err
	}

	milestonesURL := fmt.Sprintf("https://api.github.com/repos/%s/%s/milestones", owner, repo)
	resp, err = http.Get(milestonesURL)
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&milestones); err != nil {
		return nil, nil, err
	}

	return issues, milestones, nil
}

func handler(w http.ResponseWriter, r *http.Request) {
	issues, milestones, err := fetchGitHubData(os.Args[1], os.Args[2])
	if err != nil {
		log.Printf("4.14: %v\n", err)
		return
	}

	tmpl, err := template.New("index").Parse(`
		<h1>Bug Reports</h1>
		<ul>
		{{range .Issues}}
			<li><a href='{{.HTMLURL}}'>{{.Number}}</a>: {{.Title}} by <a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></li>
		{{end}}
		</ul>

		<h1>Milestones</h1>
		<ul>
		{{range .Milestones}}
			<li><a href='{{.HTMLURL}}'>{{.Number}}</a>: {{.Title}}</li>
		{{end}}
		</ul>
	`)

	if err != nil {
		log.Printf("4.14: %v\n", err)
		return
	}

	data := struct {
		Issues    []Issue
		Milestones []Milestone
	}{
		Issues:    issues,
		Milestones: milestones,
	}

	tmpl.Execute(w, data)
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
