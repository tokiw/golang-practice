package intset

import (
	"math/rand"
	"testing"
)

func TestAddAll(t *testing.T) {
	var x IntSet
	y := NewMapIntSet()
	x.AddAll(1, 144, 9, 42)
	y.AddAll(1, 144, 9, 42)

	xElems := x.Elems()
	yElems := y.Elems()
	for i := 0; i < len(xElems); i++ {
		if xElems[i] != yElems[i] {
			t.Errorf("expected %v but got %v", xElems[i], yElems[i])
		}
	}
}

func TestIntersectWith(t *testing.T) {
	var x, x2, y, y2 IntSet
	x.AddAll(1, 144, 9, 42)
	x2.AddAll(1, 2, 144, 300, 500)

	x.IntersectWith(&x2)

	y.AddAll(1, 144, 9, 42)
	y2.AddAll(1, 2, 144, 300, 500)

	y.IntersectWith(&y2)

	xElems := x.Elems()
	yElems := y.Elems()
	for i := 0; i < len(xElems); i++ {
		if xElems[i] != yElems[i] {
			t.Errorf("expected %v but got %v", xElems[i], yElems[i])
		}
	}
}

func TestDifferenceWith(t *testing.T) {
	var x, x2, y, y2 IntSet
	x.AddAll(1, 144, 9, 42)
	x2.AddAll(1, 9, 144)

	x.DifferenceWith(&x2)

	y.AddAll(1, 144, 9, 42)
	y2.AddAll(1, 9, 144)

	y.DifferenceWith(&y2)

	xElems := x.Elems()
	yElems := y.Elems()
	for i := 0; i < len(xElems); i++ {
		if xElems[i] != yElems[i] {
			t.Errorf("expected %v but got %v", xElems[i], yElems[i])
		}
	}
}

func TestSymmetricDifference(t *testing.T) {
	var x, x2, y, y2 IntSet
	x.AddAll(1, 144, 9, 42)
	x2.AddAll(1, 9)

	x.SymmetricDifference(&x2)

	y.AddAll(1, 144, 9, 42)
	y2.AddAll(1, 9)

	y.SymmetricDifference(&y2)

	xElems := x.Elems()
	yElems := y.Elems()
	for i := 0; i < len(xElems); i++ {
		if xElems[i] != yElems[i] {
			t.Errorf("expected %v but got %v", xElems[i], yElems[i])
		}
	}
}

// TODO: 乱数を外側で作って同じデータでベンチマークを取る必要がある。。。
func BenchmarkMapIntSetAdd10(b *testing.B) {
	n := 10

	for i := 0; i < b.N; i++ {
		mapintset := NewMapIntSet()
		for j := 0; j < n; j++ {
			mapintset.Add(rand.Intn(100))
		}
	}
}
func BenchmarkIntSetAdd10(b *testing.B) {
	n := 10

	for i := 0; i < b.N; i++ {
		var intset IntSet
		for j := 0; j < n; j++ {
			intset.Add(rand.Intn(100))
		}
	}
}
func BenchmarkMapIntSetAdd100(b *testing.B) {
	n := 100

	for i := 0; i < b.N; i++ {
		mapintset := NewMapIntSet()
		for j := 0; j < n; j++ {
			mapintset.Add(rand.Intn(100))
		}
	}
}
func BenchmarkIntSetAdd100(b *testing.B) {
	n := 100

	for i := 0; i < b.N; i++ {
		var intset IntSet
		for j := 0; j < n; j++ {
			intset.Add(rand.Intn(100))
		}
	}
}
func BenchmarkMapIntSetAdd1000(b *testing.B) {
	n := 1000

	for i := 0; i < b.N; i++ {
		mapintset := NewMapIntSet()
		for j := 0; j < n; j++ {
			mapintset.Add(rand.Intn(100))
		}
	}
}
func BenchmarkIntSetAdd1000(b *testing.B) {
	n := 1000

	for i := 0; i < b.N; i++ {
		var intset IntSet
		for j := 0; j < n; j++ {
			intset.Add(rand.Intn(100))
		}
	}
}
