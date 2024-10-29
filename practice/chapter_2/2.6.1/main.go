package main

import (
        "fmt"
        "os"
)

var pc [256]byte

func init() {
	for i := range pc {
			pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCount(x uint64) int {
	sum := 0
	for x > 0 {
		x = x&(x-1)
		sum++
	}
	return sum
}

func main() {
	x := 0b11111111
	fmt.Println(PopCount(0xFFFF), int(0xFFFF))
	fmt.Println(PopCount(0xFFF0), int(0xFFF0))
	fmt.Println(PopCount(0xFF00), int(0xFF00))
	fmt.Println(PopCount(0xF000), int(0xF000))
	fmt.Printf("%b, %b,", x, x&(x-1))


	os.Exit(0)
}