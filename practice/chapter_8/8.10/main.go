// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 243.

// Crawl3 crawls web links starting with the command-line arguments.
//
// This version uses bounded parallelism.
// For simplicity, it does not address the termination problem.
package main

import (
	"fmt"
	"log"
	"os"
	"sync"

	"practice/chapter_5/links"
)

var cancel = make(chan struct{})

func cancelled() bool {
	select {
	case <-cancel:
		return true
	default:
		return false
	}
}

func crawl(url string) []string {
	if cancelled() {
		return nil
	}
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

//!+
func main() {
	worklist := make(chan []string)  // lists of URLs, may have duplicates
	unseenLinks := make(chan string) // de-duplicated URLs
	var wg sync.WaitGroup

	go func() {
		os.Stdin.Read(make([]byte, 1))
		close(cancel)
	}()

	// Add command-line arguments to worklist.
	go func() { 
		select {
		case worklist <- os.Args[1:]:
		case <-cancel: 
			return
		}
		
	}()

	// Create 20 crawler goroutines to fetch each unseen link.
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				select {
				case <-cancel:
					return 
				case link := <- unseenLinks:
					foundLinks := crawl(link)
					if foundLinks != nil {
						select {
						case worklist <- foundLinks:
						case <-cancel:
							return
						}
					}
				}
			}
		}()
	}

	// The main goroutine de-duplicates worklist items
	// and sends the unseen ones to the crawlers.
	seen := make(map[string]bool)
	for {
		select {
		case list := <-worklist:
			for _, link := range list {
				if !seen[link] {
					seen[link] = true
					select {
					case unseenLinks <- link:
					case <-cancel:
						return
					}
				}
			}
		case <-cancel:
			close(worklist)
			close(unseenLinks)
			wg.Wait()         
			return
		}
	}
}

//!-
