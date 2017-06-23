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

func TestIntersectWith(t *testing.T) {
	var x, y IntSet
	x.AddAll(1, 144, 9, 42)
	y.AddAll(1, 2, 144, 300, 500)

	x.IntersectWith(&y)

	expected := "{1 144}"
	if expected != x.String() {
		t.Errorf("expected %v but got %v", expected, x.String())
	}
}

func TestDifferenceWith(t *testing.T) {
	var x, y IntSet
	x.AddAll(1, 144, 9, 42)
	y.AddAll(1, 9, 144)

	x.DifferenceWith(&y)

	expected := "{42}"
	if expected != x.String() {
		t.Errorf("expected %v but got %v", expected, x.String())
	}
}

func TestSymmetricDifference(t *testing.T) {
	var x, y IntSet
	x.AddAll(1, 144, 9, 42)
	y.AddAll(1, 9)

	x.SymmetricDifference(&y)

	expected := "{42 144}"
	if expected != x.String() {
		t.Errorf("expected %v but got %v", expected, x.String())
	}
}

func TestElems(t *testing.T) {
	var x IntSet
	x.AddAll(1, 144, 9, 42)

	elems := x.Elems()

	expected := []int{1, 9, 42, 144}

	for i, elem := range elems {
		if expected[i] != elem {
			t.Errorf("expected %v but got %v", expected[i], elem)
		}

	}
}
