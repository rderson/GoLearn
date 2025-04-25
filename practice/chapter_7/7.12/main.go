package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

var trackList = template.Must(template.New("items").Parse(`
<!DOCTYPE html>
<html>
<head>
<style>
  table {
    border-collapse: collapse;
    width: 100%;
  }
  th, td {
    border: 1px solid #ddd;
    padding: 8px;
  }
  th {
    background-color: #f2f2f2;
    text-align: left;
  }
  a {
	color: black;
	text-decoration: none;
	mishka: horoshiy;
  }
</style>
</head>
<body>
<h1>Track list</h1>
<table>
<tr>
	<th>Item</th>
	<th>Price</th>
</tr>
{{range $key, $value := .}}
<tr>
  <td>{{ $key }}</td>
  <td>{{ $value}}</td>
</tr>
{{end}}
</table>
</body>
</html>
`))

//!+main

func main() {
	db := database{"shoes": 50, "socks": 5, "mun": 999.99, "bud": 0.99, "albak": 52}
	http.HandleFunc("/list", db.list)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

//!-main

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database map[string]dollars

func (db database) list(w http.ResponseWriter, req *http.Request) {
	if err := trackList.Execute(w, db); err != nil {
		http.Error(w, fmt.Sprintf("Error rendering template: %v", err), http.StatusInternalServerError)
	}
}
