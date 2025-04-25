package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strings"
)

type Element struct {
	Name       string
	Attributes map[string]string
}

func main() {
	dec := xml.NewDecoder(os.Stdin)
	var stack []Element

	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "7.17: %v\n", err)
			os.Exit(1)
		}
		switch tok := tok.(type){
		case xml.StartElement:
			attrMap := make(map[string]string)
			for _, attr := range tok.Attr {
				attrMap[attr.Name.Local] = attr.Value
			}
			stack = append(stack, Element{Name: tok.Name.Local, Attributes: attrMap})
		case xml.EndElement:
			stack = stack[:len(stack)-1]
		case xml.CharData:
			if match(stack, os.Args[1:]) {
				data := strings.TrimSpace(string(tok))
				if data != "" {
					fmt.Printf("%s: %s\n", formatPath(stack), data)
				}
			}
		}
	}
}

func match(stack []Element,  selectors []string) bool {
	if len(stack) < len(selectors) {
		return false
	}

	for i := 0; i < len(selectors); i++ {
		selector := selectors[i]
		element := stack[len(stack)-len(selectors)+i]

		// Check if the selector matches the element name or attributes
		if !strings.Contains(selector, "=") {
			// Match by element name
			if element.Name != selector {
				return false
			}
		} else {
			// Match by attribute (e.g., id=value or class=value)
			parts := strings.SplitN(selector, "=", 2)
			if len(parts) != 2 {
				return false
			}
			attrName, attrValue := parts[0], parts[1]
			if element.Attributes[attrName] != attrValue {
				return false
			}
		}
	}

	return true
}

func formatPath(stack []Element) string {
	var path []string
	for _, el := range stack {
		path = append(path, el.Name)
	}
	return strings.Join(path, " > ")
}

