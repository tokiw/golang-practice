package popcount

import "testing"

func TestPopCpunt(t *testing.T) {
	assertEqual(0, 0, t)
	assertEqual(1, 1, t)
	assertEqual(2, 1, t)
	assertEqual(3, 2, t)
	assertEqual(4, 1, t)
	assertEqual(5, 2, t)
	assertEqual(6, 2, t)
	assertEqual(7, 3, t)
}

func assertEqual(input uint64, expected int, t *testing.T) {
	actual := PopCount(input)
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func BenchmarkPopCountTable1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountTable(uint64(1))
	}
}

func BenchmarkPopCount1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(uint64(1))
	}
}

func BenchmarkPopCountTable100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountTable(uint64(100))
	}
}

func BenchmarkPopCount100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(uint64(100))
	}
}

func BenchmarkPopCountTable1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountTable(uint64(1000))
	}
}

func BenchmarkPopCount1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(uint64(1000))
	}
}

func BenchmarkPopCountTable10000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountTable(uint64(10000))
	}
}

func BenchmarkPopCount10000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(uint64(10000))
	}
}
