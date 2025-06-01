package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"practice/chapter_12/sexpr"
	"strings"
)

type response2 struct {
    Page   int      `json:"page"`
    Fruits []string `json:"fruits"`
}

type Person struct {
	Name	string
	Age		int
}

func main() {
	str := `{"page": 1, "fruits": ["apple", "peach"]}`

	dec := json.NewDecoder(strings.NewReader(str))
	res1 := response2{}
	dec.Decode(&res1)
	fmt.Println(res1)

	sexprs:= []string{`((Name "Alice") (Age 30))`, `((Name "Bob") (Age 8))`, `((Name "Alexander") (Age 52))`}

	for _, sexp := range sexprs {
		sDec := sexpr.NewDecoder(strings.NewReader(sexp))
		res2 := Person{}
		if err := sDec.Decode(&res2); err != nil {
			fmt.Fprintf(os.Stderr, "12.7: %v\n", err)
			os.Exit(1)
		}
		fmt.Println(res2)
	}

	file, err := os.Open("decode_test.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "12.7: %v\n", err)
		os.Exit(1)
	}

	sDec := sexpr.NewDecoder(file)
	for {
		var p Person 
		if err := sDec.Decode(&p); err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "12.7: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("%s, %d\n", p.Name, p.Age)
	}

}