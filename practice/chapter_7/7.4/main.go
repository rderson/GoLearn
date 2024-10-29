package main

import (
	"fmt"
	"io"
	"os"

	"golang.org/x/net/html"
)

// создание структуры myReader с которой мы будем имплементировать собственную функцию NewReader
type myReader struct {
	s string 	// строка которую мы читаем
	pos int 	// позиция в строке на которой мы остановились (0 по умолчанию)
}

// имплементация функции NewReader которая возвращает myReader
func NewReader(s string) *myReader {
	return &myReader{s: s}
}

// имплементация метода Read для того чтобы интерфейс io.Reader был доволен. Читает до len(p) байтов со строки, затем возвращает количество прочитанных байтов, а также ошибку io.EOF в случае если мы дочитали до конца строки.
func (r *myReader) Read(p []byte) (n int, err error) {
	// дошли до конца строки, вернули EOF
	if r.pos >= len(r.s) {
		return 0, io.EOF
	}

	n = copy(p, r.s[r.pos: ]) // копируем строку в p (срез байтов)
	r.pos += n // продвигаем pos на кол-во прочитанных байтов (в некст раз Read начнет с этой позиции)

	return n, nil // возвращаем кол-во прочитанных байтов
}

func main()  {
	htmlCode := "<html><body><h1>Hello, World!</h1></body></html>" // пример строки содержащей html код

	doc, err := html.Parse(NewReader(htmlCode)) // читает html страницу из переданной нами в NewReader строки, возвращает *Node (саму страницу) и err (ошибку, в случае ее наличия)

	// чек ошибки
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}

	elements := make(map[string]int) // мапа в которой перечисляются все теги html элементов и их количество на странице 
	populateElements(elements, doc)

	// перебор и принт мапы в формате tag: quantity
	for k, v := range elements{ 
		fmt.Printf("%s: %d\n", k, v)
	}
}

func populateElements(elements map[string]int, n *html.Node)  {
	// обновляем мапу (n.Data = название тега) 
	if n.Type == html.ElementNode { 
		elements[n.Data]++
	}

	// рекурсия шобы пройтись по всем чилдренам, а затем родственным элементам
	for c := n.FirstChild; c != nil; c = c.NextSibling { 
		populateElements(elements, c)
	}
}