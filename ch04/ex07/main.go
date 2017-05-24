package main

import (
	"bytes"
	"fmt"
)

func main() {
	alpha := []byte("abcde")
	alpha = reverse(alpha)
	fmt.Println(string(alpha))
	ja := []byte("あいうえお")
	ja = reverse(ja)
	fmt.Println(string(ja))
}

func reverse(s []byte) []byte {
	runes := bytes.Runes(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return []byte(string(runes))
}
