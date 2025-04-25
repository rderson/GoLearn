package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)


func main()  {
	url := os.Args[1]

	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "5.17: %v", err)
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "5.17: %v", err)
	}

	for _, node := range ElementsByTagName(doc, "h1", "p", "img") {
		fmt.Println(node.Type, node.Attr, node.Data)
	}
}

func ElementsByTagName(doc *html.Node, names ...string) []*html.Node {
	var result []*html.Node

	if doc.Type == html.ElementNode && contains(names, doc.Data) {
		result = append(result, doc)
	}
	
	for c := doc.FirstChild; c != nil; c = c.NextSibling {
		result = append(result, ElementsByTagName(c, names...)...)
	}

	return result
}

func contains(names []string, name string) bool {
	for _, n := range names {
		if n == name {
			return true
		}
	}
	return false
}

