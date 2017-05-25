package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed file open")
	}
	input := bufio.NewScanner(file)
	input.Split(bufio.ScanWords)

	words := make(map[string]int)

	for input.Scan() {
		word := input.Text()
		words[word]++
	}

	for word, count := range words {
		fmt.Printf("%s: %d\n", word, count)
	}
}
