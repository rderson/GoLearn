package main

import (
	"fmt"
	"time"

	// "golang.org/x/tools/go/analysis/passes/timeformat"
)

func main() {
	// timeformat
	fmt.Println(time.Now().Format("02-01-2006"))
}