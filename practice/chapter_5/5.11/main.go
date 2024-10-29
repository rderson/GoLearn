// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 136.

// The toposort program prints the nodes of a DAG in topological order.
package main

import (
	"errors"
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

func addPrerequisite(course string, reqs []string)  {
	prereqs[course] = reqs
}

// !+main
func main() {
	addPrerequisite("linear algebra", []string{"calculus"})

	order, err := topoSort(prereqs)
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		for i, course := range order {
			fmt.Printf("%d:\t%s\n", i+1, course)
		}
	}
}

func topoSort(m map[string][]string) ([]string, error) {
	var order []string
	seen := make(map[string]bool)
	stack := make(map[string]bool)
	var visitAll func(items []string) error

	visitAll = func(items []string) error{
		for _, item := range items {
			if stack[item] {
				return errors.New("cycle detected involving " + item)
			}
			if !seen[item] {
				seen[item] = true
				stack[item] = true

				if err := visitAll(m[item]); err != nil {
					return err
				}

				stack[item] = false
				order = append(order, item)
			} else {
				fmt.Println(item)
			}
		}
		return nil
	}

	var keys []string
	for key := range m {
		keys = append(keys, key)
	}

	sort.Strings(keys)
	if err := visitAll(keys); err != nil {
		return nil, err
	}

	return order, nil
}

//!-main
