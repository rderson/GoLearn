package intset

import (
	"math/rand"
	"testing"
)

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


// BENCHMARKS
var iterNum int = 100000000

var result bool

func BenchmarkAdd(b *testing.B)  {
	var x IntSet
	for i := 0; i < b.N; i++ {
		x.Add(rand.Intn(iterNum))
	}
}

func BenchmarkHas(b *testing.B) {
	var x IntSet
	for i := 0; i < b.N; i++ {
		x.Add(i)
		result = x.Has(rand.Intn(iterNum))
	}
}

func BenchmarkUnionWith(b *testing.B)  {
	for i := 0; i < b.N; i++ {
		var x IntSet
		var y IntSet
		x.Add(rand.Intn(iterNum))
		y.Add(rand.Intn(iterNum))
		x.UnionWith(&y)
	}
}

func BenchmarkAddMap(b *testing.B)  {
	x := make(map[int]bool)
	for i := 0; i < b.N; i++ {
		x[rand.Intn(iterNum)] = true
	}
}

func BenchmarkHasMap(b *testing.B)  {
	x := make(map[int]bool)
	for i := 0; i < b.N; i++ {
		x[i] = true
		result = x[rand.Intn(iterNum)]
	}
}

func BenchmarkUnionWithMap(b *testing.B)  {
	for i := 0; i < b.N; i++ {
		x := make(map[int]bool)
		y := make(map[int]bool)
		x[rand.Intn(iterNum)] = true
		y[rand.Intn(iterNum)] = true
		for k := range y {
			if !x[k] {
				x[k] = true
			}
		}
	}
}