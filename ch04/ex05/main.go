package main

import (
	"fmt"
)

func main() {
	slice := []int{1, 2, 2, 3, 4, 4}
	fmt.Println(removeDuplicate(slice))
}

func removeDuplicate(s []int) []int {
	for i := 0; i < len(s)-1; {
		if s[i] == s[i+1] {
			s = remove(s, i+1)
		} else {
			i++
		}
	}
	return s
}

func remove(slice []int, i int) []int {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}
