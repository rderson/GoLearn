package main

import (
	"fmt"
	"os"
)

func main() {
	s := []int{0, 1, 2, 3, 4, 5}
	rotate(s, 2)
	fmt.Println(s)
}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func rotate(s []int, pos int)  {
	if pos > len(s) {
		fmt.Fprintln(os.Stderr, "4.4: unexpected error, rotate argument is too high")
		os.Exit(1)
	}
	reverse(s[:pos])
	reverse(s[pos:])
	reverse(s)
}