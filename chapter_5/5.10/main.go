// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 136.

// The toposort program prints the nodes of a DAG in topological order.
package main

import (
	"fmt"
	"sort"
)

// !+table
// prereqs maps computer science courses to their prerequisites.
var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},

	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},

	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

//!-table

// !+main
func main() {
	order := topoSort(prereqs)

	var keys []string
	for course := range order {
		keys = append(keys, course)
	}

	sort.Slice(keys, func(i, j int) bool {
		return order[keys[i]] < order[keys[j]]
	})

	for i, course := range keys{
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

func topoSort(m map[string][]string) map[string]int {
	order := make(map[string]int)
	seen := make(map[string]bool)
	var visitAll func(items []string)

	currentOrder := 1

	visitAll = func(items []string) {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				visitAll(m[item])
				order[item] = currentOrder
				currentOrder++
			}
		}
	}

	for key := range m {
		if !seen[key] {
			visitAll([]string{key})
		}
	}

	fmt.Println(order)
	return order
}

//!-main
