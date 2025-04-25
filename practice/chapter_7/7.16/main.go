package main

import (
	"practice/chapter_7/eval"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func parseAndCheck(s string) (eval.Expr, error) {
	if s == "" {
		return nil, fmt.Errorf("empty expression")
	}
	expr, err := eval.Parse(s)
	if err != nil {
		return nil, err
	}
	vars := make(map[eval.Var]bool)
	if err := expr.Check(vars); err != nil {
		return nil, err
	}
	for v := range vars {
		if v != "x" && v != "y" {
			return nil, fmt.Errorf("undefined variable: %s", v)
		}
	}
	return expr, nil
}

func calculator(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Error parsing form: "+err.Error(), http.StatusBadRequest)
		return
	}

	expr, err := parseAndCheck(r.Form.Get("expr"))
	if err != nil {
		http.Error(w, "Bad expr: "+err.Error(), http.StatusBadRequest)
		return
	}

	x := r.FormValue("x")
	y := r.FormValue("y")

	xParsed, err := strconv.ParseFloat(x, 64)
	if err != nil {
		http.Error(w, "Invalid value for x", http.StatusBadRequest)
		return
	}

	yParsed, err := strconv.ParseFloat(y, 64)
	if err != nil {
		http.Error(w, "Invalid value for y", http.StatusBadRequest)
		return
	}

	result := expr.Eval(eval.Env{"x": xParsed, "y": yParsed})

	fmt.Fprint(w, result)
}

func main() {
	http.HandleFunc("/calculator", calculator)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

