package main

import (
	"fmt"
	"practice/chapter_2/2.5/tempconv"
)

func main()  {
	fmt.Println("Absolute zero in Kelvin: ", tempconv.CToK(tempconv.AbsoluteZeroC))
	fmt.Println("Freezing in Kelvin: ", tempconv.CToK(tempconv.FreezingC))
	fmt.Println("Boiling in Kelvin: ", tempconv.CToK(tempconv.BoilingC))
	fmt.Println("In conclusion: who the fuck even uses Kelvin...")
}