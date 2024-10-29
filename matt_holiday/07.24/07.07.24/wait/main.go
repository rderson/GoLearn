package main

import (
	"log"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	log.Println("потерпи немного, бро...")
	time.Sleep(47 * time.Second)
	log.Println("терпилла47")
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
