package main

import (
	"fmt"
	"log"
	"practice/chapter_13/13.1/equalNum"
)

func main() {
	fmt.Println("Exercise 13.1: Define a deep comparison function that con siders numbers (of any type) equal if they differ by less than one part in a billion.")
	fmt.Println()
	

	x, y := 21.000000000000000001, 21.0000000000000005
	t, err := equalNum.Equal(x, y)
	if err != nil {
		log.Fatalf("13.1: %v", err)
	}

	fmt.Printf("%f, %f: %v", x, y, t)
}
