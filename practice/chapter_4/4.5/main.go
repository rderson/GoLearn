package main

import "fmt"

func main() {
	s := []string{"man", "man", "woman", "child", "child", "child", "dog", "dog", "dog", "dog", "cat", "cat"}
	fmt.Printf("%q\n", s)
	s = removeAdjacentDublicates(s)
	fmt.Printf("%q\n", s)
}

func remove(slice []string, i int) []string {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}

func areDublicates(slice []string) (int, bool) {
	for i := range slice {
		if i != len(slice) - 1 {
			if slice[i] == slice[i+1] {
				return i, true
			}
		}
	}
	return 0, false
}

func removeAdjacentDublicates(slice []string) []string{
	for {
		n, ok := areDublicates(slice)
		if ok {
			slice = remove(slice, n)
		} else {
			break
		}
	}
	return slice
}