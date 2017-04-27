package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}

func echo(slice []string) {
	s, sep := "", ""
	for _, arg := range slice {
		s += sep + arg
		sep = " "
	}
}

func echo1(slice []string) {
	strings.Join(slice, " ")
}
