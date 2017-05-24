package main

import "testing"

func TestComma(t *testing.T) {

	assertEqual([]int{1, 2, 3, 4}, removeDuplicate([]int{1, 2, 2, 3, 4, 4}), t)
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
