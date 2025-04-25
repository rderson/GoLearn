package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	db := database{"albak": 52, "socks": 5, "hat": 20, "skates": 49.99}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	http.HandleFunc("/new", db.new)
	http.HandleFunc("/update", db.update)
	http.HandleFunc("/remove", db.remove)
	fmt.Println("The server is running at http://localhost:8000")
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database map[string]dollars

func (db database) list(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if price, ok := db[item]; ok {
		fmt.Fprintf(w, "%s\n", price)
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Item does not exist: %s\n", item)
	}
}

func (db database) new(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price := req.URL.Query().Get("price")

	priceConv, err := strconv.ParseFloat(price, 64)
	if err != nil {
		fmt.Fprintf(w, "Error occured: %v\n", err)
	}

	if _, ok := db[item]; ok {
		fmt.Fprintf(w, "Item already exists: %s\n", item)
	} else {
		db[item] = dollars(priceConv)
		fmt.Fprintf(w, "The item %q has been added to the database.\n", item)
	}
}

func (db database) update(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price := req.URL.Query().Get("price")

	priceConv, err := strconv.ParseFloat(price, 64)
	if err != nil {
		fmt.Fprintf(w, "Error occured: %v\n", err)
	}

	if _, ok := db[item]; ok {
		db[item] = dollars(priceConv)
		fmt.Fprintf(w, "The item %q was updated.\n", item)
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Item does not exist: %s\n", item)
	}
}

func (db database) remove(w http.ResponseWriter, req *http.Request)  {
	item := req.URL.Query().Get("item")

	if _, ok := db[item]; ok {
		delete(db, item)
		fmt.Fprintf(w, "Item %q was sucessfully deleted from the database\n", item)
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Item does not exist: %s\n", item)
	}
}