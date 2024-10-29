package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// s := "хуй"

	// fmt.Printf("%8T %[1]v %d\n", s, len(s))
	// fmt.Printf("%8T %[1]v\n", []rune(s))
	// fmt.Printf("%8T %[1]v %d\n", []byte(s), len([]byte(s)))
	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "not enough args")
		os.Exit(-1)
	}

	old, new := os.Args[1], os.Args[2]
	scan := bufio.NewScanner(os.Stdin)

	for scan.Scan() {
		s := strings.Split(scan.Text(), old)
		t := strings.Join(s, new)

		fmt.Println(t)
	}
}