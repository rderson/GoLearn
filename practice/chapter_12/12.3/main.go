package main

import (
	"log"
	"practice/chapter_12/sexpr"
)

type Person struct {
	Firstname string
	Lastname  string
	Age       float64
	Hours     map[string]int
	Pidor     bool
	Hz     	  complex64
	Shiza	  interface{}
}

func main() {
	rasul := Person{
		Firstname: "Rasul",
		Lastname:  "Torshkhoev",
		Age:       17.1,
		Hours:     map[string]int{"Dota": 60, "Rock Legacy 2": 200, "AC": 10, "Spacewar": 1000},
		Pidor:     false,
		Hz: 3 + 9i,
		Shiza: []int{52, 1488},
	}
	data, err := sexpr.Marshal(rasul, false)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%s\n", data)


}