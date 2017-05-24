package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", comma(os.Args[i]))
	}
}

//!+
// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
	dot := strings.Index(s, ".")

	n := len(s)
	if n <= 3 || (dot <= 3 && dot >= 0) {
		return s
	}

	if s[0] == '-' && (n <= 4 || (dot <= 4 && dot >= 0)) {
		return s
	}

	var buf bytes.Buffer
	if dot >= 0 {
		n -= (len(s) - dot)
	}
	end := n % 3
	if end == 0 {
		end = 3
	}
	buf.WriteString(s[:end])

	start := end
	for end = start + 3; end < n && end < dot; {
		buf.WriteString("," + s[start:end])
		start, end = end, end+3
	}
	buf.WriteString("," + s[start:])
	return buf.String()
}
