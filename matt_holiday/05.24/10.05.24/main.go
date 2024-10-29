package main

import(
	"fmt"
)

func doRandomShi(b []int) (int, int) {
	b[0], b[2] = b[2], b[0]

	fmt.Printf("b@ %p\n", b)

	return 5, 2
}

func main()  {
	defer fmt.Println("52 всем нашим")

	a := []int{6, 0, 0}

	fmt.Printf("a@ %p\n", a)
	
	c, d := doRandomShi(a)

	fmt.Printf("%v %d%d\n", a, c, d)
}