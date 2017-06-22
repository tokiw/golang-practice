package intset

import (
	"testing"
)

func TestAddAll(t *testing.T) {
	var x IntSet
	x.AddAll(1, 144, 9, 42)

	expected := "{1 9 42 144}"
	if expected != x.String() {
		t.Errorf("expected %v but got %v", expected, x.String())
	}
}
