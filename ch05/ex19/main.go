package main

import (
	"fmt"
)

func main() {
	fmt.Println(returnPanic())
}
func returnPanic() (r int) {
	defer func() {
		p := recover()
		if p != nil {
			r = 1
		}
	}()

	panic("panic")
}
