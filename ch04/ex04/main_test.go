package main

import "testing"

func TestComma(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5}
	rotate(slice)
	assertEqual([]int{5, 1, 2, 3, 4}, slice, t)
}

func assertEqual(expected, actual []int, t *testing.T) {
	if len(expected) != len(actual) {
		t.Errorf("got %v\nwant %v", actual, expected)
		return
	}
	for i := 0; i < len(expected); i++ {
		if expected[i] != actual[i] {
			t.Errorf("got %v\nwant %v", actual, expected)
			break
		}
	}
}
