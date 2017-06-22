package main

import (
	"bufio"
	"bytes"
	"fmt"
)

type WordCounter int
type LineCounter int

func (c *WordCounter) Write(p []byte) (int, error) {
	s := bufio.NewScanner(bytes.NewBuffer(p))
	s.Split(bufio.ScanWords)
	var wc int
	for s.Scan() {
		wc++
	}
	*c += WordCounter(wc)
	return wc, nil
}

func (c *LineCounter) Write(p []byte) (int, error) {
	s := bufio.NewScanner(bytes.NewBuffer(p))
	var lc int
	for s.Scan() {
		lc++
	}
	*c += LineCounter(lc)
	return lc, nil
}

func main() {
	var wc WordCounter
	wc.Write([]byte("hello world"))
	fmt.Println(wc)

	var lc LineCounter
	lc.Write([]byte("hello\nworld"))
	fmt.Println(lc)
}
