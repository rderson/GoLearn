// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 101.

// Package Treesort provides insertion sort using an unbalanced binary Tree.
package treesort

import "strconv"

// !+
type Tree struct {
	value       int
	left, right *Tree
}

// Sort sorts values in place.
func Sort(values []int) {
	var root *Tree
	for _, v := range values {
		root = Add(root, v)
	}
	appendValues(values[:0], root)
}

// appendValues appends the elements of t to values in order
// and returns the resulting slice.
func appendValues(values []int, t *Tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func Add(t *Tree, value int) *Tree {
	if t == nil {
		// Equivalent to return &Tree{value: value}.
		t = new(Tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = Add(t.left, value)
	} else {
		t.right = Add(t.right, value)
	}
	return t
}

func (t *Tree) String() string {
	var values []int
	values = appendValues(values, t)

	s := ""
	for i, value := range values {
		if i != len(values) - 1 {
			s += strconv.Itoa(value) + " --> " 
		} else {
			s += strconv.Itoa(value)
		}
	} 
	return s
}

//!-
