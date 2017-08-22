package split

import (
	"strings"
	"testing"
)

func TestSplit(t *testing.T) {
	tests := []struct {
		s    string
		sep  string
		want int
	}{
		{"", "", 0},
		{"", ":", 1},
		{"a:", ":", 2},
		{":a", ":", 2},
		{"a:b", ":", 2},
		{"a:b:c", ":", 3},
	}

	for _, test := range tests {
		words := strings.Split(test.s, test.sep)
		if got := len(words); got != test.want {
			t.Errorf("Split(%q, %q) returned %d words, want %d",
				test.s, test.sep, got, test.want)
		}
	}
}
