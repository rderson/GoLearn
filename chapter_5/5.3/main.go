package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"

	"chapter_5/textformat"
)

func main()  {
	doc, err := html.Parse(os.Stdin) // читает html страницу из stdin возвращает *Node (саму страницу) и err (ошибку)

	// чек ошибки
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}

	var data = textformat.ReadAndFormatTextElements("", doc)
	fmt.Println(data)
}

