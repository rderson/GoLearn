package main

import "fmt"

type Point struct{X, Y float64}

func (p Point) Add (q Point) Point { return Point{p.X + q.X, p.Y + q.Y} }
func (p Point) Sub (q Point) Point { return Point{p.X - q.X, p.Y - q.Y} }

type Path []Point

func (path Path) TranslateBy(offset Point, add bool) {
	var op func(p, q Point) Point
	if add {
		op = Point.Add
	} else {
		op = Point.Sub
	}
	for i := range path {
		path[i] = op(path[i], offset)
	}
}

func main() {
	p := Path{{5, 4}, {3, 52}, {1488, 23}, {0, 0}, {0, 6}}
	fmt.Println("Befote translate: ", p)
	p.TranslateBy(Point{1, 1}, false)
	fmt.Println("After translate: ", p)

	fmt.Println("Ебанные долбаебы не поймут этой хуйни")
}
