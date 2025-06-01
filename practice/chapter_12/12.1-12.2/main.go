package main

import (
	"fmt"

	"practice/chapter_12/display"
)

type Shtuka struct {
	size int
	form string
}

func main() {
	m := make(map[string]int)
	m["sosi"] = 15
	m["mun"] = 52
	display.Display("m", m)
	fmt.Println()

	rasulShtuka := Shtuka{
		size: 14,
		form: "straight",
	}

	jsShtuka := Shtuka{
		size: 19,
		form: "charismatic",
	}

	display.Display("rasulShtuka", rasulShtuka)
	fmt.Println()

	pns := make(map[Shtuka]string)
	pns[rasulShtuka] = "Rasul"
	pns[jsShtuka] = "J. S."
	display.Display("pns", pns)
	fmt.Println()

	var arr = [5]int{52, 26, 1488, 1889, 1945}
	display.Display("arr", arr)
	fmt.Println()

	arrs := make(map[[5]int]string)
	arrs[arr] = "dates"
	arrs[[5]int{1, 2, 3, 4, 5}] = "natural numbers"
	display.Display("arrs", arrs)
}