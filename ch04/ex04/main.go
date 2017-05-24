package main

import (
	"fmt"
)

func main() {
	slice := []int{1, 2, 3, 4, 5}
	rotate(slice)
	fmt.Println(slice)
}

func rotate(s []int) {
	end := len(s) - 1

	for i := 0; i < len(s); i++ {
		s[i], s[end] = s[end], s[i]
	}
}
