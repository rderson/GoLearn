package main

import (
	"strings"
	"testing"
)

func TestSplit(t *testing.T) {
	var tests = []struct {
		s 		string
		sep 	string
		want 	int
	}{
		{"a:b:c", ":", 3},
		{"sosi hui", " ", 2},
		{"pachimu, pachimu, pachimu, pachimu... ai balya, pohui", ",", 5},
		{"g;n;i;d;a;<3", ";", 6},
		{"s0u0k0a", "0", 4},
	}
	for _, test := range tests {
		words := strings.Split(test.s, test.sep)
		if got, want := len(words), test.want; got != want {
			t.Errorf("Split(%q, %q) returned %d words, want %d", test.s, test.sep, got, want)
		}
	}
	
}
