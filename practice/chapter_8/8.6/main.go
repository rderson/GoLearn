package main

import (
	"fmt"
	"log"
	"os"

	"practice/chapter_5/links"
)

var depth = 3

func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

// !+
func main() {
	worklist := make(chan []string)  // lists of URLs, may have duplicates
	unseenLinks := make(chan string) // de-duplicated URLs

	var n, d int 

	n++
	counter := make([]int, depth+2)
	counter[d] = n
	// Add command-line arguments to worklist.
	go func() { worklist <- os.Args[1:] }()

	// Create 20 crawler goroutines to fetch each unseen link.
	for i := 0; i < 20; i++ {
		go func() {
			for link := range unseenLinks {
				foundLinks := crawl(link)
				go func() { worklist <- foundLinks }()
			}
		}()
	}

	// The main goroutine de-duplicates worklist items
	// and sends the unseen ones to the crawlers.
	seen := make(map[string]bool)
	for ; n > 0; n-- {

		list := <- worklist

		if d > depth {
			continue
		}

		for _, link := range list {
			if !seen[link] {
				n++
				counter[d+1]++
				seen[link] = true
				unseenLinks <- link
			}
		}

		if counter[d]--; counter[d] == 0 {
			d ++
		}
	}
}

//!-
