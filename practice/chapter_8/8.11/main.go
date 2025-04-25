package main

import (
	"fmt"
	"log"
	"net/http"
)

var done = make(chan struct{})

func mirroredQuery(query []string) string {
	responses := make(chan string, len(query))
	
	for _, hostname := range query {
		go func(hostname string)  {
			select {
			case responses <- hostname + ": " + request(hostname):
				close(done)
			case <-done:
				return
			}
			
		}(hostname)
	}

	return <-responses // return the quickest respons
}

func request(hostname string) (string) {
	url := "https://" + hostname
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("8.11: %v", err)
		return ""
	}
	defer resp.Body.Close()
	return resp.Status
}

func main() {
	query := []string{"google.com", "ozon.ru", "amazon.de", "bund.de", "welt.de"}
	fmt.Println(mirroredQuery(query))
}
