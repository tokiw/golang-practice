package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {
	counts := make(map[rune]int)    // counts of Unicode characters
	var utflen [utf8.UTFMax + 1]int // count of lengths of UTF-8 encodings
	invalid := 0                    // count of invalid UTF-8 characters

	in := bufio.NewReader(os.Stdin)
	var typeCounts map[string]int
	for {
		r, n, err := in.ReadRune() // returns rune, nbytes, error
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		counts[r]++
		utflen[n]++
		typeCounts = countChars(r)
	}
	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Print("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	fmt.Print("\ntype\tcount\n")
	for c, n := range typeCounts {
		fmt.Printf("%s\t%d\n", c, n)
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}

func countChars(c rune) map[string]int {
	var counts = make(map[string]int)

	if unicode.IsLetter(c) {
		counts["letter"]++
	}
	if unicode.IsControl(c) {
		counts["control"]++
	}
	if unicode.IsDigit(c) {
		counts["digit"]++
	}
	if unicode.IsGraphic(c) {
		counts["graphic"]++
	}
	if unicode.IsLower(c) {
		counts["lower"]++
	}
	if unicode.IsMark(c) {
		counts["mark"]++
	}
	if unicode.IsNumber(c) {
		counts["number"]++
	}
	if unicode.IsPrint(c) {
		counts["print"]++
	}
	if !unicode.IsPrint(c) {
		counts["noprint"]++
	}
	if unicode.IsPunct(c) {
		counts["punct"]++
	}
	if unicode.IsSpace(c) {
		counts["space"]++
	}
	if unicode.IsSymbol(c) {
		counts["symbol"]++
	}
	if unicode.IsTitle(c) {
		counts["title"]++
	}
	if unicode.IsUpper(c) {
		counts["upper"]++
	}
	return counts
}
