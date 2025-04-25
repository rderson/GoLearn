package main

import (
	"practice/chapter_7/treesort"
	"fmt"
	"math/rand"
)

func main() {
	data := make([]int, 50)
	for i := range data {
		data[i] = rand.Int() % 50
	}
	fmt.Println(data)
	treesort.Sort(data)
	fmt.Println(data)

	var ass *treesort.Tree

	ass = treesort.Add(ass, 5)
	ass = treesort.Add(ass, 10)
	ass = treesort.Add(ass, 1)
	ass = treesort.Add(ass, 15)
	ass = treesort.Add(ass, 35)

	fmt.Printf("Tree: %s\n", ass.String())
}