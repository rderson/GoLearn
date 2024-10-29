package main

import "fmt"

func min(vals ...int) int{
	minVal := vals[0] 
	for _, v := range vals {
		if v < minVal {
			minVal = v
		}
	}
	return minVal
}

func max(vals ...int) int{
	maxVal := vals[0]
	for _, v := range vals {
		if v > maxVal {
			maxVal = v
		}
	}
	return maxVal
}

func main() {
	s := []int{1, 125, 1098, 2, -52, 10, 510, 1488}
	fmt.Println(min(1, 2, 3, 4, 5), max(1, 2, 3, 4, 5))
	fmt.Println(min(s...), max(s...))
}