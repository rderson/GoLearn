package main

import(
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "Hi, bitch, from %s!\n", r.URL.Path[1:])	
}

func main()  {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}