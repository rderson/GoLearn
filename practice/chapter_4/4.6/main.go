package main

import (
	"fmt"
	"unicode"
)

func main() {
	input := []byte("This \n\n is \t\n a              kokojambo,\n\n   you   know?")
	fmt.Println(string(input))
	input = cleanUp(input)
	fmt.Println(string(input))
}

func cleanUp(b []byte) []byte {
	b = replaceSpaces(b)
	b = removeAdjacentSlashes(b)
	b = replaceSlashes(b)
	return b
}

func replaceSpaces(b []byte) []byte {
	var noSpaces []byte
	for _, v := range b {
		if !unicode.IsSpace(rune(v)) {
			noSpaces = append(noSpaces, v)
		} else {
			noSpaces = append(noSpaces, byte('/'))
		}
	}
	return noSpaces
}

func replaceSlashes(b []byte) []byte {
	var noSlash []byte
	for _, v := range b {
		if rune(v) != '/' {
			noSlash = append(noSlash, v)
		} else {
			noSlash = append(noSlash, byte(' '))
		}
	}
	return noSlash
}

func remove(slice []byte, i int) []byte {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}

func areDublicateSlashes(slice []byte) (int, bool) {
	for i := range slice {
		if i != len(slice) - 1 {
			if slice[i] == byte('/') && slice[i+1] == byte('/') {
				return i, true
			}
		}
	}
	return 0, false
}

func removeAdjacentSlashes(slice []byte) []byte{
	for {
		n, ok := areDublicateSlashes(slice)
		if ok {
			slice = remove(slice, n)
		} else {
			break
		}
	}
	return slice
}
