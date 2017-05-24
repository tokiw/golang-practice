package main

import "testing"

func TestComma(t *testing.T) {
	assertEqual(true, anagram("a", "a"), t)
	assertEqual(true, anagram("ab", "ab"), t)
	assertEqual(true, anagram("ab", "ba"), t)
	assertEqual(true, anagram("ab c", "a cb"), t)
}

func assertEqual(expected, actual bool, t *testing.T) {
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}
