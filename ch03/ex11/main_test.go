package main

import "testing"

func TestComma(t *testing.T) {
	assertEqual("123", comma("123"), t)
	assertEqual("1,234", comma("1234"), t)
	assertEqual("123.4", comma("123.4"), t)
	assertEqual("1,234.5", comma("1234.5"), t)
	assertEqual("-123", comma("-123"), t)
	assertEqual("-1,234", comma("-1234"), t)
	assertEqual("-123.4", comma("-123.4"), t)
	assertEqual("-1,234.5", comma("-1234.5"), t)
}

func assertEqual(expected, actual string, t *testing.T) {
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}
