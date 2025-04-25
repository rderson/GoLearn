package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strings"
)

type Node interface{}

type CharData string

type Element struct {
	Type     xml.Name
	Attr     []xml.Attr
	Children []Node
}

func main() {
	file, err := os.Open("kaif.xml")
	if err != nil {
		fmt.Fprintf(os.Stderr, "7.18: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	dec := xml.NewDecoder(file)
	root, err := buildTree(dec)
	if err != nil {
		fmt.Fprintf(os.Stderr, "7.18: %v\n", err)
		os.Exit(1)
	}

	printTree(root, 0)
}

func buildTree(dec *xml.Decoder) (Node, error) {
	var stack []*Element
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}

		switch tok := tok.(type) {
		case xml.StartElement:
			el := &Element{
				Type: tok.Name,
				Attr: tok.Attr,
			}
			if len(stack) > 0 {
				parent := stack[len(stack)-1]
				parent.Children = append(parent.Children, el)
			}
			stack = append(stack, el)
		case xml.EndElement:
			if len(stack) == 0 {
				return nil, fmt.Errorf("unexpected end of element: %s", tok.Name.Local)
			}
			el := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				return el, nil
			}
		case xml.CharData:
			if len(stack) > 0 {
				text := string(tok)
				text = trimWhitespace(text)
				if text != "" {
					parent := stack[len(stack)-1]
					parent.Children = append(parent.Children, CharData(text))
				}
			}
		}
	}
	return nil, fmt.Errorf("end of XML")
}

func printTree(node Node, depth int) {
	space := ""
	for i := 0; i < depth; i++ {
		space += "\t"
	}

	switch node := node.(type) {
	case *Element:
		fmt.Printf("%s<%s>\n", space, node.Type.Local)
		for _, attr := range node.Attr {
			fmt.Printf("%s\tAttribute: %s=\"%s\"\n", space, attr.Name.Local, attr.Value)
		}
		for _, child := range node.Children {
			printTree(child, depth+1)
		}
		fmt.Printf("%s</%s>\n", space, node.Type.Local)
	case CharData:
		text := string(node)
		text = trimWhitespace(text) // Удаляем лишние пробелы и символы переноса
		if text != "" {             // Выводим только непустой текст
			fmt.Printf("%s%s\n", space, text)
		}
	default:
		fmt.Printf("%sUnknown node type\n", space)
	}
}

func trimWhitespace(s string) string {
	return strings.TrimSpace(s)
}

