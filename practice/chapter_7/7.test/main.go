package main

import (
	"fmt"
	"io"
	"os"
	"sort"
)

var w io.Writer
var x io.Writer

type Penis struct{
	Owner string
	Size int
}

type PenisSlice []Penis

func (p PenisSlice) Len() int				{ return len(p) }
func (p PenisSlice) Less(i, j int) bool		{ return p[i].Size < p[j].Size }
func (p PenisSlice) Swap(i, j int)			{ p[i], p[j] = p[j], p[i] }

func main()  {
	w = os.Stdout
	w.Write([]byte("yoyoyo Jessie Pinkman in da haus!\n"))

	x = os.Stdout
	if w == x {
		x.Write([]byte("idi gulay popoi vilay\n"))
	} else {
		x.Write([]byte("wattaheeeeell omagad\n"))
	}

	s := []string{"Yekaterinburg", "Adelaide", "Donetsk", "California", "Budapest"}

	sort.Strings(s)

	fmt.Println(s)

	pencils := PenisSlice{{"Mun_1", 15}, {"Mun_2", 1}, {"Mun_3", 12000}, {"Mun_4", 4}}

	for _, pencil := range pencils {
		fmt.Printf("%s - %d\n", pencil.Owner, pencil.Size)
	}

	fmt.Println("-----")

	sort.Sort(pencils)

	for _, pencil := range pencils {
		fmt.Printf("%s - %d\n", pencil.Owner, pencil.Size)
	}
}