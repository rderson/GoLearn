// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 58.
//!+

// Surface computes an SVG rendering of a 3-D surface function.
package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"os"
	"strconv"
)

func main() {
	usage := "usage: main.go saddle|eggbox|paraboloid|cone|baza|antibaza web(optional)"
	var f fTypeFunc
	if len(os.Args) < 2 || len(os.Args) > 3{
		fmt.Fprintln(os.Stderr, usage)
		os.Exit(1)
	}
	switch os.Args[1] {
	case "saddle":
		f = saddle
	case "eggbox":
		f = eggbox
	case "paraboloid":
		f = paraboloid
	case "cone":
		f = cone
	case "baza":
		f = baza
	case "antibaza":
		f = antibaza
	default:
		fmt.Fprintln(os.Stderr, usage)
		os.Exit(1)
	}
	if len(os.Args) > 2 && os.Args[2] == "web"{
		handler := func(w http.ResponseWriter, r *http.Request)  {
			w.Header().Set("Content-Type", "image/svg+xml")
			if err := r.ParseForm(); err != nil {
				log.Print(err)
			}
			for k, values := range r.Form {
				for _, v := range values {
					if k == "width" {
						var err error
						width, err = strconv.Atoi(v)
						if err != nil {
							log.Println(err)
							os.Exit(1)
						}
					}
					if k == "height" {
						var err error
						height, err = strconv.Atoi(v)
						if err != nil {
							log.Println(err)
							os.Exit(1)
						}
					}
				}
			}
			svg(w, f)
		}
		http.HandleFunc("/", handler)
		log.Fatal(http.ListenAndServe("localhost:8000", nil))
		return
	}
	svg(os.Stdout, f)
}

var (
	width, height = 1200, 700           // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

type fTypeFunc func(x, y float64) float64

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°

func svg(w io.Writer, f fTypeFunc) {

	zmin, zmax := minmax(f)
	fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7; margin: auto;' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j, f)
			bx, by := corner(i, j, f)
			cx, cy := corner(i, j+1, f)
			dx, dy := corner(i+1, j+1, f)
			if math.IsNaN(ax) || math.IsNaN(ay) || math.IsNaN(bx) || math.IsNaN(by) || math.IsNaN(cx) || math.IsNaN(cy) || math.IsNaN(dx) || math.IsNaN(dy) {
				continue
			}
			fmt.Fprintf(w, "<polygon style='stroke: %s;' points='%g,%g %g,%g %g,%g %g,%g'/>\n", color(i, j, zmin, zmax, f),
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Fprintln(w, "</svg>")
}

func minmax(f fTypeFunc) (min float64, max float64) {
	min = math.NaN()
	max = math.NaN()
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			for xoff := 0; xoff <= 1; xoff++ {
				for yoff := 0; yoff <= 1; yoff++ {
					x := xyrange * (float64(i+xoff)/float64(cells) - 0.5)
					y := xyrange * (float64(j+yoff)/float64(cells) - 0.5)
					z := f(x, y)
					if math.IsNaN(min) || z < min {
						min = z
					}
					if math.IsNaN(max) || z > max {
						max = z
					}
				}
			}
		}
	}
	return
}

func color(i, j int, zmin, zmax float64, f fTypeFunc) string {
	min := math.NaN()
	max := math.NaN()
	for xoff := 0; xoff <= 1; xoff++ {
		for yoff := 0; yoff <= 1; yoff++ {
			x := xyrange * (float64(i+xoff)/float64(cells) - 0.5)
			y := xyrange * (float64(j+yoff)/float64(cells) - 0.5)
			z := f(x, y)
			if math.IsNaN(min) || z < min {
				min = z
			}
			if math.IsNaN(max) || z > max {
				max = z
			}
		}
	}

	color := ""
	if math.Abs(max) > math.Abs(min) {
		red := math.Exp(math.Abs(max)) / math.Exp(math.Abs(zmax)) * 255
		if red > 255 {
			red = 255
		}
		color = fmt.Sprintf("#%02x0000", int(red))
	} else {
		blue := math.Exp(math.Abs(min)) / math.Exp(math.Abs(zmin)) * 255
		if blue > 255 {
			blue = 255
		}
		color = fmt.Sprintf("#0000%02x", int(blue))
	}
	return color
}


func corner(i, j int, f fTypeFunc) (sx, sy float64) {
	var (
		xyscale       = float64(width) / 2 / xyrange // pixels per x or y unit
		zscale        = float64(height) * 0.4        // pixels per z unit
	)
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/float64(cells) - 0.5)
	y := xyrange * (float64(j)/float64(cells)- 0.5)

	// Compute surface height z.
	z := f(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx = float64(width)/2 + (x-y)*cos30*xyscale
	sy = float64(height)/2 + (x+y)*sin30*xyscale - z*zscale
	return
}

func eggbox(x, y float64) float64 {
	return 0.2 * (math.Cos(x) + math.Cos(y))
}

func saddle(x, y float64) float64 {
	a := 25.0
	b := 17.0
	a2 := a * a
	b2 := b * b
	return (y*y/a2 - x*x/b2)
}

func paraboloid(x, y float64) float64 {
	x2 := x * x
	y2 := y * y
	return (3*x2 + 5*y2) / 500
}

func cone(x, y float64) float64 {
	x2 := x * x
	y2 := y * y
	return math.Sqrt(x2 + y2)*0.1
}

func baza(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}

func antibaza(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Tan(r) / r
}


//
