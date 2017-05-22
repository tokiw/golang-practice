package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", comma(os.Args[i]))
	}
}

//!+
// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}

	var buf bytes.Buffer
	end := n % 3

	if end == 0 {
		end = 3
	}
	buf.WriteString(s[:end])

	start := end
	for end = start + 3; end < n; {
		buf.WriteString("," + s[start:end])
		start, end = end, end+3
	}
	buf.WriteString("," + s[start:])
	return buf.String()
}

//!-
