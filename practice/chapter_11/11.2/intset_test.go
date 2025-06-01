package intset

import "testing"

func TestIntSet(t *testing.T) {
	var x IntSet
	y := make(map[int]bool)
	x.Add(4)
	x.Add(5)
	x.Add(52)
	y[4], y[5], y[52] = true, true, true
	
	if !x.Has(52) || !y[52] {
		t.Error("x.Has(52) = false or y[52] = false")
	}

	var z IntSet
	values := []int{1488, 77, 123}

	for _, v := range values {
		z.Add(v)
		y[v] = true
	}

	x.UnionWith(&z)

	for i := range y {
		if !x.Has(i) {
			t.Error("!x.Has(i)")
		}
	}
}

