package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

var mapping = map[string]string{"a": "href", "img": "src", "script": "src", "link": "href"}

func main() {
	doc, err := html.Parse(os.Stdin) // читает html страницу из stdin возвращает *Node (саму страницу) и err (ошибку)

	// чек ошибки
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}

	for _, link := range visit("a", nil, doc) {
		fmt.Println(link)
	}

	fmt.Println("--------------------------------------------------")

	for _, link := range visit("img", nil, doc) {
		fmt.Println(link)
	}

	fmt.Println("--------------------------------------------------")

	for _, link := range visit("script", nil, doc) {
		fmt.Println(link)
	}

	fmt.Println("--------------------------------------------------")

	for _, link := range visit("link", nil, doc) {
		fmt.Println(link)
	}
}

func visit(target string, links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == target {
		for _, a := range n.Attr {
			if a.Key == mapping[target] {
				links = append(links, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(target, links, c)
	}
	return links
}