package main

import (
	"fmt"
	"log"
	"net/http"
	"practice/chapter_12/params"
)


func searchValid(resp http.ResponseWriter, req *http.Request) {
	var data struct {
		Email     	string 		`http:"email" valid:"email"`
		MaxResults 	int      	`http:"max"`
		Exact      	bool     	`http:"x"`
		Zip 		int			`http:"zip" valid:"zip"`
	}
	if err := params.Unpack(req, &data); err != nil {
		http.Error(resp, err.Error(), http.StatusBadRequest) // 400
		return
	}

	// ...rest of handler...
	fmt.Fprintf(resp, "Search: %+v\n", data)
}

//!-

func main() {
	http.HandleFunc("/search", searchValid)
	log.Fatal(http.ListenAndServe(":12345", nil))
}
