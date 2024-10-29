package main

import(
	"fmt"
	"math"
)

type Point struct {
	X, Y float64
}

type Line struct {
	Begin, End Point
}

type Path []Point

func (l Line) Distance() float64 {
	return math.Hypot(l.End.X - l.Begin.X, l.End.Y - l.Begin.Y)
}

func (p Path) Distance() (sum float64) {
	for i := 1; i < len(p); i++{
		sum += Line{p[i-1], p[i]}.Distance()
	}
	return sum
}

type Distancer interface  {
	Distance() float64
}

func PrintDistance(d Distancer){
	fmt.Println(d.Distance())
}

func main() {
	side := Line{Point{1, 2}, Point{4, 6}}
	perimeter := Path{{1, 1}, {5, 1}, {5, 4}, {1, 1}}

	PrintDistance(side)
	PrintDistance(perimeter)
}