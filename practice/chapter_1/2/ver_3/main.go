package main

import(
	"fmt"
	"os"
	"strings"
)

func main()  {
	counts := make(map[string]int)

	for _, file := range os.Args[1:] {
		data, err := os.ReadFile(file)
		if err != nil {
			fmt.Fprintf(os.Stderr, "ver_3: %v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			if line != "\r" {
				counts[line]++
			}
		}
	}
	
	for line, n := range counts {
		if n > 1{
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}