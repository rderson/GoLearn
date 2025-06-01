package main

import (
	"encoding/json"
	"fmt"
	"log"
	"practice/chapter_12/sexpr"
)

type Developer struct {
	Name    string
	Country string
}

type VideoGame struct {
	Title       string                 // string
	Year        int                    // int
	Price       float64                // float
	Multiplayer bool                   // bool
	Tags        []string               // slice
	Metadata    map[string]string      // map       		
	Developer   *Developer             // pointer to struct
	Extras      interface{}            // interface{}
}

func main() {
	var p *int
	f := 25.25
	b := true

	fmt.Printf("%v\n", p)
	fmt.Printf("%v\n", f)
	fmt.Printf("%v\n", b)

	data, err := json.Marshal(p)
	if err != nil{
		log.Fatal(err)
	}

	fmt.Printf("%s\n", data)

	data, err = json.Marshal(f)
	if err != nil{
		log.Fatal(err)
	}

	fmt.Printf("%s\n", data)

	data, err = json.Marshal(b)
	if err != nil{
		log.Fatal(err)
	}

	fmt.Printf("%s\n", data)

	dev := &Developer{
		Name:    "Bethesda Game Studios",
		Country: "USA",
	}

	game := VideoGame{
		Title:       "Skyrim",
		Year:        2011,
		Price:       39.99,
		Multiplayer: false,
		Tags:        []string{"RPG", "Open World", "Fantasy"},
		Metadata:    map[string]string{"engine": "Creation Engine", "rating": "M"},
		Developer:   dev,
		Extras:      []interface{}{"DLCs included", "100+ hours", true},
	}

	data, err = sexpr.Marshal(game, true)
	if err != nil {
		log.Fatalf("Marshal failed: %v", err)
	}

	fmt.Printf("%s\n", data)

	var decoded VideoGame
	err = json.Unmarshal(data, &decoded)
	if err != nil {
		log.Fatalf("Unmarshal failed: %v", err)
	}

	fmt.Println("\n🔹 Unmarshaled Struct:")
	fmt.Printf("%+v\n", decoded)

	fmt.Println("\n🔹 Developer Name:", decoded.Developer.Name)
	fmt.Println("🔹 Extras (raw):", decoded.Extras)
}
