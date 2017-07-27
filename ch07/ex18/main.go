package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strings"
)

type Node interface{} // CharData or *Element
type CharData string
type Element struct {
	Type     xml.Name
	Attr     []xml.Attr
	Children []Node
}

func main() {
	dec := xml.NewDecoder(os.Stdin)
	var root *Element
	var stack []*Element // stack of element names
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "xmlselect: %v\n", err)
			os.Exit(1)
		}
		switch tok := tok.(type) {
		case xml.StartElement:
			var element Element
			element.Type = tok.Name
			element.Attr = tok.Attr
			if root == nil {
				root = &element
			} else {
				parent := stack[len(stack)-1]
				parent.Children = append(parent.Children, &element)
			}
			stack = append(stack, &element) // push
		case xml.EndElement:
			stack = stack[:len(stack)-1] // pop
		case xml.CharData:
			if containsAll(stack, os.Args[1:]) {
				fmt.Printf("%s: %s\n", strings.Join(stack, " "), tok)
			}
		}
	}
}
