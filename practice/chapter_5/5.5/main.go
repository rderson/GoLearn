package main

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/html"

	"practice/chapter_5/textformat"
)

var imgCount = 0

func main()  {
	doc, err := html.Parse(os.Stdin) // читает html страницу из stdin возвращает *Node (саму страницу) и err (ошибку)

	// чек ошибки
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}

	var words, images = countWordsAndImages(doc)
	fmt.Printf("Words: %d\nImages: %d\n", words, images)
}

func countImages(n *html.Node) {
	if n.Type == html.ElementNode && n.Data == "img" {
		imgCount++
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		countImages(c)
	}
}

func countWordsAndImages(n *html.Node) (wordCount, imageCount int) {
	data := textformat.ReadAndFormatTextElements("", n)
	words := strings.Fields(data)
	wordCount = len(words)
	countImages(n)
	imageCount = imgCount
	return
}