// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 91.

//!+nonempty

// Nonempty is an example of an in-place slice algorithm.
package main

import "fmt"

// nonempty returns a slice holding only the non-empty strings.
// The underlying array is modified during the call.
func nonempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}

//!-nonempty

func remove(slice []int, i int) []int {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}

func main() {
	//!+main
	data := []string{"one", "", "three"}
	fmt.Printf("%q\n", nonempty(data)) // `["one" "three"]`
	fmt.Printf("%q\n", data)           // `["one" "three" "three"]`
	//!-main
	familie := []string{"Toriel", "", "", "Asgor", "", "Asriel", "", "", "", "Chara"}
	fmt.Print("Home: ")
	familie = nonempty(familie)
	for i, s := range familie {
		if i != len(familie) - 1 {
			fmt.Printf("%s, ", s)
		} else {
			fmt.Printf("%s\n", s)
		}
	}
	z := []int{0, 1, 2, 256, 3, 148, 292, 4, 5, 926}
	z = remove(z, 3)
	z = remove(z, 4)
	z = remove(z, 4)
	z = remove(z, 6)
	fmt.Println(z)
}

// !+alt
func nonempty2(strings []string) []string {
	out := strings[:0] // zero-length slice of original
	for _, s := range strings {
		if s != "" {
			out = append(out, s)
		}
	}
	return out
}

//!-alt
