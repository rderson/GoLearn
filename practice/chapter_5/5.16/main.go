package main

import (
	"fmt"
	"strings"
)

func joinANALog(sep string, elems ...string) string {
	var before []string
	var after string
	before = append(before, elems...)
	after = strings.Join(before, sep)
	return after
}

func main() {
	s1 := "Ja"
	s2 := "sosu"
	s3 := "4orniy"
	s4 := "hui!"
	fmt.Println(joinANALog(" ", s1, s2, s3, s4))
}