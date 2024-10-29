package main

import (
	"fmt"
)

type Pair struct {
	Path string
	Hash string
}

func (p Pair) String() string {
	return fmt.Sprintf("Hash of %s is %s", p.Path, p.Hash)
}

type PairWithLength struct {
	Pair
	Length int
}

func (pwl PairWithLength) String() string{
	return fmt.Sprintf("Hash of %s is %s and the length is %d", pwl.Path, pwl.Hash, pwl.Length)
}

func main() {
	p := Pair{"/usr", "0xfdfe"}
	pwl := PairWithLength{Pair{"/usr/lol", "0xdead"}, 133}

	fmt.Println(p)
	fmt.Println(pwl)
}