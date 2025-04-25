package main

import (
	"fmt"
	"sort"
)

type Palindrome []byte

func (p Palindrome) Len() int           { return len(p) }
func (p Palindrome) Less(i, j int) bool { return p[i] < p[j] }
func (p Palindrome) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

var pon Palindrome

func main() {
	pon = []byte("sosi")
	fmt.Println(isPalindrome(pon))
	pon = []byte("sosu")	
	fmt.Println(isPalindrome(pon))
	pon = []byte("soos")
	fmt.Println(isPalindrome(pon))
}

func isPalindrome(s sort.Interface) bool {
	n := s.Len()
	for i := 0; i < n/2; i++ {
		j := n - 1 - i
		if s.Less(i, j) || s.Less(j, i) {
			return false
		}
	}
	
	return true
}