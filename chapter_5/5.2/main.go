package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main()  {
	doc, err := html.Parse(os.Stdin) // читает html страницу из stdin возвращает *Node (саму страницу) и err (ошибку)

	// чек ошибки
	if err != nil {
		fmt.Fprintf(os.Stderr, "5.2: %v\n", err)
		os.Exit(1)
	}

	elements := make(map[string]int)
	populateElements(elements, doc)

	for k, v := range elements{
		fmt.Printf("%s: %d\n", k, v)
	}
}

func populateElements(elements map[string]int, n *html.Node)  {
	if n.Type == html.ElementNode {
		elements[n.Data]++
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		populateElements(elements, c)
	}
}