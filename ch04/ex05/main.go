package main

import (
	"fmt"
)

func main() {
	slice := []string{"1", "2", "2", "3", "4", "4"}
	fmt.Println(removeDuplicate(slice))
}

func removeDuplicate(s []string) []string {
	for i := 0; i < len(s)-1; {
		if s[i] == s[i+1] {
			s = remove(s, i+1)
		} else {
			i++
		}
	}
	return s
}

func remove(slice []string, i int) []string {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}
