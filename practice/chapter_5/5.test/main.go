package main

import (
	"practice/chapter_5/links"
	"fmt"
)

func main() {
	allLinks, _ := links.Extract("https://google.com")
	for _, link := range allLinks {
		fmt.Println(link)
	}
}