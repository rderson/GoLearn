package main

import(
	"fmt"
	"strconv"
	"strings"
)

type IntSlice []int

func (iS IntSlice) String() string {
	var strs []string

	for _, v := range iS{
		strs = append(strs, strconv.Itoa(v))
	}

	return "[" + strings.Join(strs, " GOIDA ") + "]" 
}

func main()  {
	var slicer IntSlice = []int{1, 2, 3}
	var slicerCopy fmt.Stringer = slicer

	for i, v := range slicer{
		fmt.Printf("%d: %d\n", i, v)
	}

	fmt.Printf("%T %[1]v\n", slicer)
	fmt.Printf("%T %[1]v\n", slicerCopy)
}