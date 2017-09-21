package circular

import (
	"fmt"
)

func Example_IsCircular() {
	fmt.Println(IsCircular(1.0)) // "false"

	// struct
	type sample struct {
		Inner interface{}
	}
	var obj sample
	obj.Inner = obj

	fmt.Println(IsCircular(obj)) // "true"

	// map
	m := make(map[string]interface{})
	m["key"] = m
	fmt.Println(IsCircular(obj)) // "true"

	// slice
	var s []interface{}
	s = append(s, s)
	fmt.Println(IsCircular(s)) // "true"

	// Output:
	// false
	// true
	// true
	// true
}
