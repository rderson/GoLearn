package main

import (
	"fmt"
	"math"
	"os"
	"strconv"

	"practice/chapter_2/2.6/conv"
)

func main()  {
	for _, arg := range os.Args[2:] {
		v, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "conv: %v\n", err)
			os.Exit(1)
		}

		if os.Args[1] == "temp" {
			f := conv.Fahrenheit(v)
			c := conv.Celsius(v)
			fmt.Printf("%s = %s, %s = %s\n", f, conv.FToC(f), c, conv.CToF(c))
		} else if os.Args[1] == "length" {
			m := conv.Meters(v)
			ft := conv.Feet(v)
			fmt.Printf("%s = %s, %s = %s\n", m, conv.Feet(math.Round(float64(conv.MToFt(m)*100))/100), ft, conv.Meters(math.Round(float64(conv.FtToM(ft)*100))/100))
		} else if os.Args[1] == "weight" {
			kg := conv.Kilos(v)
			p := conv.Pounds(v)
			fmt.Printf("%s = %s, %s = %s\n", kg, conv.Pounds(math.Round(float64(conv.KToP(kg)*100))/100), p, conv.Kilos(math.Round(float64(conv.PToK(p)*100))/100))
		} else {
			fmt.Fprintln(os.Stderr, "Enter one of the following arguments into console: 'temp', 'length', 'weight'.")
			os.Exit(1)
		}
	}
}