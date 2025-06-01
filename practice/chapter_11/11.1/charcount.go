// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 97.
//!+

// Charcount computes counts of Unicode characters.
package charcount

import (
	"io"
	"strings"
	"unicode"
	"unicode/utf8"
)

func Charcount(input string) (counts map[rune]int, utflen [utf8.UTFMax + 1]int, invalid int, err error) {
	counts = make(map[rune]int)
	in := strings.NewReader(input)
	for {
		r, n, err := in.ReadRune() // returns rune, nbytes, error
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, [5]int{0, 0, 0, 0, 0}, 0, err
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		counts[r]++
		utflen[n]++
	}
	return
}

//!-
