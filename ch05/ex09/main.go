package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	fmt.Println(expand("ab$cde", func(str string) string {
		return strings.ToUpper(str)
	}))
	fmt.Println(expand("ab$cde$fg", func(str string) string {
		runes := []rune(str)
		for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
			runes[i], runes[j] = runes[j], runes[i]
		}
		return string(runes)
	}))
	fmt.Println(expand("あい$うえお$かき", func(str string) string {
		runes := []rune(str)
		for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
			runes[i], runes[j] = runes[j], runes[i]
		}
		return string(runes)
	}))
}

func expand(s string, f func(string) string) string {
	r := regexp.MustCompile(`\$[^$]+`)
	return r.ReplaceAllStringFunc(s, func(s string) string { return f(s[1:]) })
}
