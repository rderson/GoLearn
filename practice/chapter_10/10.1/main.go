// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 287.

//!+main

// The jpeg command reads a PNG image from the standard input
// and writes it as a JPEG image to the standard output.
package main

import (
	"errors"
	"flag"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png" // register PNG decoder
	"io"
	"os"
)

func main() {
	outputFormat := flag.String("f", "jpg", "output format")
	flag.Parse()

	if err := toFormat(os.Stdin, os.Stdout, *outputFormat); err != nil {
		fmt.Fprintf(os.Stderr, "jpeg: %v\n", err)
		os.Exit(1)
	}
}

func toFormat(in io.Reader, out io.Writer, format string) error {
	img, kind, err := image.Decode(in)
	if err != nil {
		return err
	}
	fmt.Fprintln(os.Stderr, "Input format =", kind)
	switch format {
	case "gif":
		return gif.Encode(out, img, &gif.Options{NumColors: 256})
	case "png":
		return png.Encode(out, img)
	case "jpg":
		return jpeg.Encode(out, img, &jpeg.Options{Quality: 95})
	default:
		return errors.New("unsupported format")
	}
}

//!-main

/*
//!+with
$ go build gopl.io/ch3/mandelbrot
$ go build gopl.io/ch10/jpeg
$ ./mandelbrot | ./jpeg >mandelbrot.jpg
Input format = png
//!-with

//!+without
$ go build gopl.io/ch10/jpeg
$ ./mandelbrot | ./jpeg >mandelbrot.jpg
jpeg: image: unknown format
//!-without
*/
