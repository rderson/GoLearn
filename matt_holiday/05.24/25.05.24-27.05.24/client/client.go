package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

const url = "https://jsonplaceholder.typicode.com"

func main()  {
	resp, err := http.Get(url + "/todos/1")

	type todo struct {
		ID 			int 	`json:"id"`
		Title 		string 	`json:"title"`
		Completed 	bool 	`json:"completed"`
	}

	if err != nil{
		fmt.Fprintln(os.Stderr, err)
		os.Exit(-1)
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		body, err := io.ReadAll(resp.Body)

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(-2)
		}

		var item todo

		err = json.Unmarshal(body, &item)

		if err != nil{
			fmt.Fprintln(os.Stderr, err)
			os.Exit(-3)
		}

		fmt.Printf("%#v\n", item)
	}
}