package main

import (
	"practice/chapter_7/eval"
	"fmt"
)

func main() {
	expr, err := eval.Parse("sin(pow(x, 2)-pow(y,2))/10")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Expression: ", expr.String())
	if expr.String() == eval.Format(expr) {
		fmt.Println("The syntax trees are equivalent!")
	} else {
		fmt.Println("The syntax trees are NOT equivalent!")
	}
}
