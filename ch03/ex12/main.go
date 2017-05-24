package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if anagram(os.Args[1], os.Args[2]) {
		fmt.Println("Anagram!!")
	} else {
		fmt.Println("Not Anagram")
	}
}

func anagram(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i := 0; i < len(s1); i++ {
		if !strings.Contains(s2, string(s1[i])) {
			return false
		}
	}
	return true
}
