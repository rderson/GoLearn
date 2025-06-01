package charcount

import "testing"

func TestCharcount(t *testing.T)  {
	counts, _, _, _ := Charcount("sosi penis.")
	if counts['s'] != 3 {
		t.Error(`counts['s'] != 3`)
	}

	counts, _, _, _ = Charcount("aaabbbccc")
	if counts['a'] != 3 && counts['b'] != 3 && counts['c'] != 3 {
		t.Error(`counts['a'] != 3 && counts['b'] != 3 && counts['c'] != 3`)
	}
}

func TestUTFLen(t *testing.T)  {
	_, utflen, _, _ := Charcount("sosi penis.")
	if utflen[2] != 0 {
		t.Error("utflen[2] != 0")
	}

	_, utflen, _, _ = Charcount("ХАХАХА")
	if utflen[2] != 6 {
		t.Error("utflen[2] != 6")
	}
}