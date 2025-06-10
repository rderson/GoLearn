package main

import (
	"fmt"
	"practice/chapter_13/13.2/cycle"
)

type Node struct {
	value int
	tail  *Node
}

func main() {
	a := Node{1, nil}
	b := Node{1488, nil}
	a.tail = &b
	b.tail = &a

	if cycle.IsCyclic(a) {
		fmt.Println("Cycle found!", a)
	} else {
		fmt.Println("Cycle not found.")
	}

	c := Node{52, nil}
	d := Node{26, nil}
	c.tail = &d

	if cycle.IsCyclic(c) {
		fmt.Println("Cycle found!", c)
	} else {
		fmt.Println("Cycle not found.")
	}

}