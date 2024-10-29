package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// func main()  {
// a, c := 12, 345
// fmt.Printf("%d %d\n", a, c)
// fmt.Printf("%#x %#x\n", a, c)

// fmt.Println()

// fmt.Printf("|%6d|%6d|\n", a, c)
// fmt.Printf("|%06d|%06d|\n", a, c)
// fmt.Printf("|%-6d|%-6d|\n", a, c)

// s := []int{1,2,3}
// arr := [3]rune{'a', 'b', 'c'}
// m := map[string]int{"and":1, "or": 2}
// stroke := "a stroke"
// b := []byte(stroke)

// fmt.Printf("%T\n", s)
// fmt.Printf("%v\n", s)
// fmt.Printf("%#v\n", s)
// fmt.Println()
// fmt.Printf("%T\n", arr)
// fmt.Printf("%q\n", arr)
// fmt.Printf("%v\n", arr)
// fmt.Printf("%#v\n", arr)
// fmt.Println()
// fmt.Printf("%T\n", m)
// fmt.Printf("%v\n", m)
// fmt.Printf("%#v\n", m)
// fmt.Println()
// fmt.Printf("%T\n", stroke)
// fmt.Printf("%v\n", stroke)
// fmt.Printf("%#v\n", stroke)
// fmt.Printf("%v\n", string(b))
// }
func main()  {
	// for _, fname := range os.Args[1:] {
	// 	file, err := os.Open(fname)

	// 	if err != nil{
	// 		fmt.Fprintln(os.Stderr, err)
	// 		continue
	// 	}
		
	// 	data, err := io.ReadAll(file)

	// 	if err != nil {
	// 		fmt.Fprint(os.Stderr, err)
	// 		continue
	// 	}

	// 	fmt.Println("The file has", len(data), "bytes")
	// 	file.Close()
	// }
	for _, fname := range os.Args[1:] {

		var lineCount, wordCount, charachterCount int

		file, err := os.Open(fname)

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}

		scan := bufio.NewScanner(file)

		for scan.Scan() {
			s := scan.Text()

			wordCount += len(strings.Fields(s))
			charachterCount += len(s)
			lineCount++
		}

		fmt.Printf(" %7d %7d %7d %s\n", lineCount, wordCount, charachterCount, fname)
		file.Close()
	}
}