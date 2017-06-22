package join

import (
	"testing"
)

func TestJoin(t *testing.T) {
	assertJoin(t, "", "")
	assertJoin(t, "", ".")
	assertJoin(t, "a", "", "a")
	assertJoin(t, "a", ".", "a")
	assertJoin(t, "ab", "", "a", "b")
	assertJoin(t, "a.b", ".", "a", "b")
	assertJoin(t, "abc", "", "a", "b", "c")
	assertJoin(t, "a.b.c", ".", "a", "b", "c")
	assertJoin(t, "abcd", "", "a", "b", "c", "d")
	assertJoin(t, "a.b.c.d", ".", "a", "b", "c", "d")
}

func assertJoin(t *testing.T, expected string, sep string, vals ...string) {
	if r := Join(sep, vals...); r != expected {
		t.Errorf("result: %v, expected: %v", r, expected)
	}
}
