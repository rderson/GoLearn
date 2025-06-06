package main

import (
	"fmt"
	"practice/chapter_12/params"
)

type Info struct {
	Labels     []string `http:"l"`
	MaxResults int      `http:"max"`
	Exact      bool     `http:"x"`
}

func main() {
	data := Info{
		Labels: []string{"da", "ya", "lublu", "sosat", "chlen"},
		MaxResults: 52,
		Exact: true,
	}
	fmt.Printf("Result: %q", params.Pack("http://localhost:12345/search/", data))
}