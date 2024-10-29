package main

import (
	"fmt"
	"chapter_4/4.test_2/treesort"
)

func main()  {
	values := []int{2, 512, 8, 4, 32, 128, 64, 256, 1024, 16}

	fmt.Println("Before treesort: ", values)

	treesort.Sort(values)

	fmt.Println("After treesort: ", values)
}