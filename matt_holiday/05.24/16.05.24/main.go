package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/html"
)

var raw = `<!DOCTYPE html>
<html>
  <body>
    <h1>My First Heading</h1>
      <p>My first paragraph.</p>
      <p>HTML <a href="https://www.w3schools.com/html/html_images.asp">images</a> are defined with the img tag:</p>
      <img src="xxx.jpg" width="104" height="142">
  </body>
</html>`

func visit(n *html.Node, words, pics * int){
	if n.Type == html.TextNode {
		*words += len(strings.Fields(n.Data))
	} else if n.Type == html.ElementNode && n.Data == "img" {
		*pics++
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		visit(c, words, pics)
	}
}

func countWordsAndImages(doc *html.Node) (int, int) {
	var words, pics int

	visit(doc, &words, &pics)

	return words, pics
}

func main()  {
	doc, err := html.Parse(bytes.NewReader([]byte(raw)))

	if err != nil {
		fmt.Fprintf(os.Stderr, "parse failed: %s\n", err)
		os.Exit(-1)
	}

	words, pics := countWordsAndImages(doc)

	fmt.Printf("%d words %d images", words, pics)
}