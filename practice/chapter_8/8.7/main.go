/*
Exercise 8.7: Write a concurrent program that creates a local mirror of a website, fetching
each reachable page and writing it to a directory on the local disk. Only pages with in the
original domain (for instance, golang.org) should be fetched. URLs within mirrored pages
should be altered as needed so that they refer to the mirrored page, not the original.
*/

package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"practice/chapter_5/links"

	"golang.org/x/net/html"
)

func modifyLinks(content []byte, rootDomain string) []byte  {
	doc, err := html.Parse(bytes.NewReader(content))
	if err != nil {
		log.Printf("error parsing content as html document: %v\n", err)
		return content
	}

	var modify func(n *html.Node)
	modify = func(n *html.Node)  {
		if n.Type == html.ElementNode {
			for i, attr := range n.Attr {
				if attr.Key == "href" || attr.Key == "src" {
					parsedURL, err := url.Parse(attr.Val)
					if err == nil && parsedURL.Host == rootDomain {
						localPath := parsedURL.Path
						if strings.HasSuffix(localPath, "/") || localPath == "" {
							localPath += "index.html"
						} else if filepath.Ext(localPath) == "" {
							localPath += ".html"
						}
						n.Attr[i].Val = localPath
					}
				}
			}
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			modify(c)
		}
	}
	modify(doc)

	var buf bytes.Buffer
	html.Render(&buf, doc)
	return buf.Bytes()
}

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
	
	if err := savePage(urlStr, body, rootDomain); err != nil {
		log.Printf("error saving the page: %v\n", err)
	}

	list, err := links.Extract(urlStr)
	if err != nil {
		log.Printf("error extracting the links: %v\n", err)
		return nil
	}

	return list
}

func savePage(urlStr string, data []byte, rootDomain string) error {
	parsedURL, err := url.Parse(urlStr)
	if err != nil {
		return fmt.Errorf("error parsing url: %v", err)
	}
	
	data = modifyLinks(data, rootDomain)

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

	return os.WriteFile(filePath, data, 0644)
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("usage: ./mirror <url>")
	}

	worklist := make(chan []string)  // lists of URLs, may have duplicates
	unseenLinks := make(chan string) // de-duplicated URLs

	// Add command-line arguments to worklist.
	go func() { worklist <- os.Args[1:] }()

	startURL := os.Args[1]
	
	parsedURL, err := url.Parse(startURL)
	if err != nil {
		log.Fatalf("start url parsing error: %v\n", err)
	}

	rootDomain := parsedURL.Host

	// Create 20 crawler goroutines to fetch each unseen link.
	for i := 0; i < 20; i++ {
		go func() {
			for link := range unseenLinks {
				foundLinks := crawl(link, rootDomain)
				go func() { worklist <- foundLinks }()
			}
		}()
	}

	// The main goroutine de-duplicates worklist items
	// and sends the unseen ones to the crawlers.
	seen := make(map[string]bool)
	for list := range worklist {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				unseenLinks <- link
			}
		}
	}
}