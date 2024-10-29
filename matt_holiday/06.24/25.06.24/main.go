package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type dollar float32

func (d dollar) String() string {
	return fmt.Sprintf("$%.2f", d)
}

type database map[string]dollar

func (db database) list(w http.ResponseWriter, req *http.Request)  {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func (db database) add(w http.ResponseWriter, req *http.Request)  {
	item := req.URL.Query().Get("item")
	price := req.URL.Query().Get("price")

	if _, ok := db[item]; ok {
		msg := fmt.Sprintf("duplicate item: %q", item)
		http.Error(w, msg, http.StatusBadRequest)
		return
	}

	p, err := strconv.ParseFloat(price, 32)

	if err != nil {
		msg := fmt.Sprintf("invalid price: %q", item)
		http.Error(w, msg, http.StatusBadRequest)
		return
	}

	db[item] = dollar(p)

	fmt.Fprintf(w, "added %s with the price %s\n", item, db[item])
}

func (db database) update(w http.ResponseWriter, req *http.Request)  {
	item := req.URL.Query().Get("item")
	price := req.URL.Query().Get("price")

	if _, ok := db[item]; !ok {
		msg := fmt.Sprintf("no such item: %q", item)
		http.Error(w, msg, http.StatusNotFound)
		return
	}

	p, err := strconv.ParseFloat(price, 32)

	if err != nil {
		msg := fmt.Sprintf("invalid price: %q", item)
		http.Error(w, msg, http.StatusBadRequest)
		return
	}

	db[item] = dollar(p)

	fmt.Fprintf(w, "new price %s for %s\n", db[item], item)
}

func (db database) look(w http.ResponseWriter, req *http.Request)  {
	item := req.URL.Query().Get("item")

	if _, ok := db[item]; !ok {
		msg := fmt.Sprintf("no such item: %q", item)
		http.Error(w, msg, http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "item %s has price %s\n", item, db[item])
}

func (db database) drop(w http.ResponseWriter, req *http.Request)  {
	item := req.URL.Query().Get("item")

	if _, ok := db[item]; !ok {
		msg := fmt.Sprintf("no such item: %q", item)
		http.Error(w, msg, http.StatusNotFound)
		return
	}

	delete(db, item)
	fmt.Fprintf(w, "item %s has been deleted\n", item)

}

func main()  {
	db := database{
		"shoes": 50,
		"socks": 5,
	}

	http.HandleFunc("/list", db.list)
	http.HandleFunc("/create", db.add)
	http.HandleFunc("/update", db.update)
	http.HandleFunc("/read", db.look)
	http.HandleFunc("/delete", db.drop)

	
	log.Fatal(http.ListenAndServe(":8080", nil))
}