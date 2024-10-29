package main

import(
	"fmt"
)

func main() {
	car := map[string]string{"Brand": "Audi", "Model": "SVO", "Year": "1488"}

	for k, v := range car {
		fmt.Printf("%8v: %v\n", k, v)
	}
	var (
		x, y int
		z float64
		s string
	)
	x,y = 9,1
	z = 6.1
	s = "papidzi"
	fmt.Println(x+y, z+8.0, s)
	fmt.Println("Uncle" + " " + "Bogdan" == "Uncle Bogdan")
}