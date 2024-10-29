package main

import(
	"fmt"
)

func main()  {
	items := [][2]byte{{0, 1}, {2, 3}, {4, 5}}
	a := [][]byte{}

	for _, item := range items {
		a = append(a, item[:])
	}

	fmt.Println(items)
	fmt.Println(a)
}