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
