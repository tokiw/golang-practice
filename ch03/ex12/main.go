package main

import (
	"bytes"
	"fmt"
	"os"
)

type Runes []rune

func main() {
	if anagram(os.Args[1], os.Args[2]) {
		fmt.Println("Anagram!!")
	} else {
		fmt.Println("Not Anagram")
	}
}

func anagram(s1, s2 string) bool {
	runes1 := bytes.Runes([]byte(s1))
	runes2 := bytes.Runes([]byte(s2))

	if len(runes1) != len(runes2) {
		return false
	}
	for i := 0; i < len(runes1); i++ {
		if !contains(runes2, runes1[i]) {
			return false
		}
	}
	return true
}

func contains(array []rune, item rune) bool {
	for _, r := range array {
		if item == r {
			return true
		}
	}
	return false
}
