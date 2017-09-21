package sexpr

import (
	"testing"
)

func TestUnmarshal(t *testing.T) {
	type Record struct {
		B    bool
		C128 complex128
	}
	r := Record{false, 1 - 2i}
	data, err := Marshal(r)
	if err != nil {
		t.Fatalf("return err %v", err.Error())
	}

	var record Record
	if err := Unmarshal(data, record); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	t.Logf("Unmarshal() = %+v\n", record)
}
