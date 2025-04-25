package main

import (
	"practice/chapter_7/eval"
	"fmt"
)

func main() {
	var expression string

	fmt.Print("Enter the expression: ")

	_, err := fmt.Scanln(&expression)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	var x float64
	fmt.Print("Variable X: ")

	_, err = fmt.Scanln(&x)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	var y float64
	fmt.Print("Variable Y: ")

	_, err = fmt.Scanln(&y)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	expr, err := eval.Parse(expression)
	if err != nil {
		fmt.Println("Error parsing expression: ", err)
	}

	vars := make(map[eval.Var]bool)

	if err = expr.Check(vars); err != nil {
		fmt.Println("Check error: ", err)
	}

	result := expr.Eval(eval.Env{"x": x, "y": y})
	fmt.Println("Result: ", result)
}
