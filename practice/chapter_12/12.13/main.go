package main

import (
	"fmt"
	"log"
	"practice/chapter_12/sexpr"
)


type VideoGame struct {
	Title       string                 `sexpr:"Nazvanie"`	// string
	Price       float64                `sexpr:"Money"`		// float  		
}

func main() {
	game := VideoGame{
		Title:       "Skyrim",
		Price:       39.99,
	}

	fmt.Println()
	data, err := sexpr.Marshal(game, false)
	if err != nil {
		log.Fatalf("Marshal failed: %v", err)
	}

	fmt.Printf("%s\n", data)

	fmt.Println()
	var g struct {
		Title       string                `sexpr:"Nazvanie"`					
		Price       float64                `sexpr:"Money"`	              									 								
	}
	if err := sexpr.Unmarshal(data, &g); err != nil {
		log.Fatalf("Unmarshal failed: %v", err)
		
	}
	fmt.Printf("%+v\n", g)
}
