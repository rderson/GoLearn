// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 61.
//!+

// Mandelbrot emits a PNG image of the Mandelbrot fractal.
package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"math/cmplx"
	"net/http"
	"os"
	"strconv"
)

func main() {

	if len(os.Args) > 1 && os.Args[1] == "web" {
		handler := func(w http.ResponseWriter, r *http.Request) {
			if err := r.ParseForm(); err != nil {
				log.Println(err)
			}
			if _, exists := r.Form["x_value"]; !exists {
				xmin, xmax = -2, +2
			}
			if _, exists := r.Form["y_value"]; !exists {
				ymin, ymax = -2, +2
			}
			if _, exists := r.Form["width"]; !exists {
				width=1024
			}
			if _, exists := r.Form["height"]; !exists {
				height=1024
			}
			for k, values := range r.Form {
				for _, v := range values {
					if k == "x_value" {
						var err error
						xmax, err = strconv.Atoi(v)
						if err != nil {
							log.Println(err)
							os.Exit(1)
						}
						xmin, _ = strconv.Atoi(v)
						xmin = xmin * -1
						log.Printf("xmin = %v, xmax = %v", xmin, xmax)
					}
					if k == "y_value" {
						var err error
						ymax, err = strconv.Atoi(v)
						if err != nil {
							log.Println(err)
							os.Exit(1)
						}
						ymin, _ = strconv.Atoi(v)
						ymin = ymin * -1
						log.Printf("ymin = %v, ymin = %v", ymin, ymax)
					}
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
			img := makeIMG()
			png.Encode(w, img)
		}
		http.HandleFunc("/", handler)
		log.Fatal(http.ListenAndServe("localhost:8000", nil))
		return
	}
	img := makeIMG()
	png.Encode(os.Stdout, img) // NOTE: ignoring errors
}

var (
	xmin, ymin, xmax, ymax = -2, -2, +2, +2
	width, height          = 1024, 1024
	epsX                   = (xmax - xmin) / width
	epsY                   = (ymax - ymin) / height
)

func makeIMG() *image.RGBA {
	offX := []float64{float64(-epsX), float64(epsX)}
	offY := []float64{float64(-epsY), float64(epsY)}

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/float64(height)*float64(ymax-ymin) + float64(ymin)
		for px := 0; px < width; px++ {
			x := float64(px)/float64(width)*float64(xmax-xmin) + float64(xmin)
			subPixels := make([]color.Color, 0)
			for i := 0; i < 2; i++ {
				for j := 0; j < 2; j++ {
					z := complex(x+offX[i], y+offY[j])
					subPixels = append(subPixels, mandelbrot(z))
				}
			}
			img.Set(px, py, avg(subPixels))
		}
	}

	return img
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}

func avg(colors []color.Color) color.Color {
	var r, g, b, a uint16
	n := len(colors)
	for _, c := range colors {
		r_, g_, b_, a_ := c.RGBA()
		r += uint16(r_ / uint32(n))
		g += uint16(g_ / uint32(n))
		b += uint16(b_ / uint32(n))
		a += uint16(a_ / uint32(n))
	}
	return color.RGBA64{r, g, b, a}
}
//!-
