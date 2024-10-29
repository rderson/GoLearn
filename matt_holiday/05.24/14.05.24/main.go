package main

import(
	"fmt"
)

func main(){
	var s []int

	t := []int{}
	u := make([]int, 5)
	v := make([]int, 0, 5)
	stuv := make([][]int, 4)
	stuv[0], stuv[1], stuv[2], stuv[3] = s, t, u, v

	for i := 0; i < len(stuv); i++ {
		fmt.Printf("%d, %d, %T, %5t, %#[3]v\n", len(stuv[i]), cap(stuv[i]), stuv[i], stuv[i]==nil)
	}
}