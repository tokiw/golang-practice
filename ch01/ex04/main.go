package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]map[string]struct{})
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for line, filenameMap := range counts {
		if len(filenameMap) > 0 {
			fmt.Printf("%s", line)
			for filename := range filenameMap {
				fmt.Printf("\t%s", filename)
			}
			fmt.Println()
		}
	}
}

func countLines(f *os.File, counts map[string]map[string]struct{}) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		txt := input.Text()
		if nil == counts[txt] {
			counts[txt] = make(map[string]struct{})
		}
		counts[txt][f.Name()] = struct{}{}
	}
	// NOTE: ignoring potential errors from input.Err()
}
