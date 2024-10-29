	// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
	// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

	// See page 139.

	// Findlinks3 crawls the web, starting with the URLs on the command line.
	package main

	import (
		"fmt"
		"io"
		"log"
		"net/http"
		"net/url"
		"os"
		"path/filepath"
		"strings"

		"chapter_5/links"
	)

	// !+breadthFirst
	// breadthFirst calls f for each item in the worklist.
	// Any items returned by f are added to the worklist.
	// f is called at most once for each item.
	func breadthFirst(f func(item string) []string, worklist []string) {
		seen := make(map[string]bool)
		for len(worklist) > 0 {
			items := worklist
			worklist = nil
			for _, item := range items {
				if !seen[item] {
					seen[item] = true
					worklist = append(worklist, f(item)...)
				}
			}
		}
	}

	//!-breadthFirst

	func savePage(urlStr string, content []byte) error {
		parsedURL, err := url.Parse(urlStr)
		if err != nil {
			return fmt.Errorf("error parsing url: %v", err)
		}

		filePath := filepath.Join(parsedURL.Host, parsedURL.Path)
		if parsedURL.Path == "" || strings.HasSuffix(parsedURL.Path, "/") {
			filePath = filepath.Join(filePath, "index.html")
		} else {
			if ext := filepath.Ext(filePath); ext == "" {
				filePath += ".html"
			}
		}

		if err := os.MkdirAll(filepath.Dir(filePath), 0755); err != nil {
			return fmt.Errorf("error creating directory: %v", err)
		}

		return os.WriteFile(filePath, content, 0644)
	}

	// !+crawl
	func crawl(urlStr string, rootDomain string) []string {
		parsedURL, err := url.Parse(urlStr)
		if err != nil {
			log.Printf("error parsing url: %v\n", err)
			return nil
		}

		if parsedURL.Host != rootDomain {
			log.Printf("skipping URL %s from a different domain\n", urlStr)
			return nil
		}

		resp, err := http.Get(urlStr)
		if err != nil {
			log.Printf("error getting url: %v\n", err)
			return nil
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			log.Printf("getting %s: %s\n", urlStr, resp.Status)
			return nil
		}

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Printf("error reading response body: %v\n", err)
			return nil
		}

		if err := savePage(urlStr, body); err != nil {
			log.Printf("error saving page %s: %v\n", urlStr, err)
		}

		list, err := links.Extract(urlStr)
		if err != nil {
			log.Printf("error extracting the links: %v\n", err)
			return nil
		}

		return list
	}

	//!-crawl

	// !+main
	func main() {
		if len(os.Args) < 2 {
			log.Fatal("usage: main.go <url>")
		}

		startURL := os.Args[1]
		if startURL == "https://golang.org" {
			startURL = "https://go.dev"
		}

		parsedURL, err := url.Parse(startURL)
		if err != nil {
			log.Fatalf("start url parsing error: %v\n", err)
		}

		rootDomain := parsedURL.Host

		breadthFirst(func(url string) []string {
			return crawl(url, rootDomain)
		}, os.Args[1:])
	}

	//!-main
