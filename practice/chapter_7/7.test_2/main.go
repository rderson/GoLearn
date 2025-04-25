package main

import (
	"fmt"
	"syscall"
)

func main() {
	fmt.Println("Ai balya ya togo mamu ebal")
	for i := 0; i < 999; i++ {
		fmt.Println(syscall.Errno(i))
	}
}