// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 122.
//!+main

// Findlinks1 prints the links in an HTML document read from standard input.
package main

import (
	"fmt"
	"os"
	
	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin) // читает html страницу из stdin возвращает *Node (саму страницу) и err (ошибку)

	// чек ошибки
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}

	// цикл перебирает все ссылки из среза возвращеннного функцией visit и печатает их по одной
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}

//!-main

// !+visit
// visit appends to links each link found in n and returns the result.
func visit(links []string, n *html.Node) []string {
	// если n это ссылка добавить ее в срез links
	if n.Type == html.ElementNode && n.Data == "a" {  
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}

	if n.FirstChild != nil {
		links = visit(links, n.FirstChild)
	}

	if n.NextSibling != nil {
		links = visit(links, n.NextSibling)
	}

	// на каждый дочерний элемент n ебашим рекурсию visit, а затем переходим на следующий родственный элемент
	// for c := n.FirstChild; c != nil; c = c.NextSibling {
	// 	links = visit(links, c)
	// 	counterA++
	// }
	return links
}

//!-visit

/*
//!+html
package html

type Node struct {
	Type                    NodeType
	Data                    string
	Attr                    []Attribute
	FirstChild, NextSibling *Node
}

type NodeType int32

const (
	ErrorNode NodeType = iota
	TextNode
	DocumentNode
	ElementNode
	CommentNode
	DoctypeNode
)

type Attribute struct {
	Key, Val string
}

func Parse(r io.Reader) (*Node, error)
//!-html
*/
