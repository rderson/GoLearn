package main

import (
	"chapter_6/intset"
	"fmt"
)

func main() {
	var x, y intset.IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	fmt.Println(x.String())

	y.Add(9)
	y.Add(42)
	fmt.Println(y.String())

	x.UnionWith(&y)
	fmt.Println(x.String()) // "{1 9 42 144}"
	fmt.Println(x.Len())
	z := x.Copy()

	x.Remove(9)
	x.AddAll(52, 1488, 69)
	fmt.Println(x.String())
	fmt.Println(y.String())
	x.SymmetricDifference(&y)
	fmt.Println(x.String())

	for i, elem := range x.Elems() {
		fmt.Printf("%d: %v\n", i+1, elem)
	}
	fmt.Println(x.Has(9), x.Has(123), z.Has(9))
}