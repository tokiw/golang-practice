package main

import (
	"fmt"
)

func main() {
	slice := []int{1, 2, 3, 4, 5}
	rotate(slice, 1)
	fmt.Println(slice)
}

func rotate(s []int, num int) {
	end := len(s) - 1

	for i := 0; i < num; i++ {
		for j := 0; j < len(s); j++ {
			s[j], s[end] = s[end], s[j]
		}

	}
}
