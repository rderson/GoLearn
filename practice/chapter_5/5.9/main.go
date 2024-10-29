package main

import (
	"fmt"
	"strings"
)

func main() {
	wish := "I really love $foo! All of my friends like $foo as well! I think $foo is the best thing in the world!"
	wish = expand(wish, trueWish)
	fmt.Println(wish)
}

func expand(s string, f func(string) string) string {
	substring := "$foo"
	s = strings.Replace(s, substring, f(substring), -1)
	return s
}

func trueWish(s string) string {
	if s == "$foo" {
		s = "DEMACIA"
	}
	return s
}

