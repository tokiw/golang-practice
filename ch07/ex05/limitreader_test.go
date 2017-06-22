package limitreader

import (
	"strings"
	"testing"
)

func TestLimitReader(t *testing.T) {
	input := "hello"
	limit := 5
	lr := LimitReader(strings.NewReader(input), int64(limit))
	i, err := lr.Read([]byte(input))
	if err != nil {
		t.Errorf("Read Error")
		return
	}
	if i != limit {
		t.Errorf("Result = %c, Expected %c", 2, i)
	}
}

func TestLimitReaderOver(t *testing.T) {
	input := "hello"
	limit := 3
	lr := LimitReader(strings.NewReader(input), int64(limit))
	i, err := lr.Read([]byte(input))
	if err != nil {
		t.Errorf("Read Error")
		return
	}
	if i != limit {
		t.Errorf("Result = %c, Expected %c", 2, i)
	}
}
