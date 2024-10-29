package main

import (
	"fmt"
	"os"
)

func main() {
	// a := 2
	// b := 6.1

	// fmt.Printf("a: %8T %[1]v\n", a)
	// fmt.Printf("b: %8T %[1]v\n", b)

	// a = int(b)
	// fmt.Printf("a: %8T %[1]v\n", a)
	// b = float64(a)
	// fmt.Printf("b: %8T %[1]v\n", b)
	var sum float64
	var n int

	for {
		var val float64

		_, err := fmt.Fscanln(os.Stdin, &val)
		if err != nil {
			break
		}

		sum += val
		n++
	}

	if n == 0 {
		fmt.Fprintln(os.Stderr, "no values")
		os.Exit(-1)
	}

	fmt.Println("The average is", sum/float64(n))
}