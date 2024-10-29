package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	// t := []byte("string")

	// fmt.Println(len(t), t)
	// fmt.Println(t[2])
	// fmt.Println(t[:2])
	// fmt.Println(t[2:])
	// fmt.Println(t[3:5], len(t[3:5]))

	scan := bufio.NewScanner(os.Stdin)
	words := make(map[string]int)

	scan.Split(bufio.ScanWords)
	for scan.Scan() {
		words[scan.Text()]++
	}

	fmt.Println(len(words), "unique words")

	type kv struct {
		key string
		val int
	}

	var ss []kv

	for k, v := range words {
		ss = append(ss, kv{k, v})
	}

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].val > ss[j].val
	})

	for _, s := range ss[:3]{
		if s.val == 1 {
			fmt.Println(s.key, "appears", s.val, "time")
		} else {
			fmt.Println(s.key, "appears", s.val, "times")
		}
	}
}