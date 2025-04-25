package main

import (
	"practice/chapter_7/eval"
	"fmt"
)

func main() {
	expr := eval.Minimum{
		Args: []eval.Expr{
			eval.Literal(52),
			eval.Var("x"),
			eval.Literal(1488),
		},
	}

	vars := make(map[eval.Var]bool)

	if err := expr.Check(vars); err != nil {
		fmt.Println("Check error: ", err)
	}

	fmt.Println("Variables: ", vars)

	fmt.Println("Expression: ", expr.String())

	min := expr.Eval(eval.Env{"x": 777})

	fmt.Println("Result: ", min)
}
