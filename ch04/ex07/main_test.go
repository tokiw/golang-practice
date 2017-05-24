package main

import "testing"

func TestComma(t *testing.T) {

	assertEqual([]byte("edcba"), reverse([]byte("abcde")), t)
	assertEqual([]byte("おえういあ"), reverse([]byte("あいうえお")), t)
}

func assertEqual(expected, actual []byte, t *testing.T) {
	if string(expected) != string(actual) {
		t.Errorf("got %v\nwant %v", actual, expected)

	}
}
