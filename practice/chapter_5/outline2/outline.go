// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 133.

// Outline prints the outline of an HTML document tree.
package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	for _, url := range os.Args[1:] {
		outline(url)
	}
}

func outline(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return err
	}

	//!+call
	if node := ElementByID(doc, "start"); node != nil {
		fmt.Println(node.Type, node.Data)
	}
	//!-call

	return nil
}

// !+forEachNode
// forEachNode calls the functions pre(x) and post(x) for each node
// x in the tree rooted at n. Both functions are optional.
// pre is called before the children are visited (preorder) and
// post is called after (postorder).

func forEachNode(n *html.Node, pre, post func(n *html.Node) bool) (*html.Node, bool) {
	if pre != nil {
		if pre(n) {
			return n, true
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if node, ok := forEachNode(c, pre, post); ok {
			return node, true
		}
	}

	if post != nil {
		if post(n) {
			return n, true
		}
	}

	return nil, false
}

//!-forEachNode

// !+startend
var depth int

func ElementByID(doc *html.Node, id string) *html.Node {
	pre := func(doc *html.Node) bool {
		for _, attr := range doc.Attr {
			if attr.Key == "id" && attr.Val == id {
				return true
			}
		}
		return false
	}

	if node, found := forEachNode(doc, pre, nil); found {
		return node
	}
	
	return nil
}

func startElement(n *html.Node) {
	if n.Type == html.ElementNode {
		fmt.Printf("%*s<%s", depth*2, "", n.Data)
		for _, attribute := range n.Attr{
			fmt.Printf(" %s=\"%s\"", attribute.Key, attribute.Val)
		}
		child := ""
		if n.Data == "img" && n.FirstChild == nil {
			child = " /"
		}
		fmt.Printf("%s>\n", child)
		depth++
	} else if n.Type == html.CommentNode {
		fmt.Printf("%*s<!--%s", depth*2, "", n.Data)
	} else if n.Type == html.TextNode {
		for _, line := range strings.Split(n.Data, "\n") {
			line = strings.TrimSpace(line)
			if line != "" && line != "\n" {
				fmt.Printf("%*s%s\n", depth*2, "", line)
			}
		}
	}
}

func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		depth--
		fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
	} else if n.Type == html.CommentNode {
		fmt.Println("-->")
	}
}

//!-startend
