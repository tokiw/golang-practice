package sexpr

import (
	"testing"
)

func TestComplex64(t *testing.T) {
	var tests = []struct {
		c    complex64
		want string
	}{
		{1 + 2i, "#C(1.000000 2.000000)"},
	}
	for _, test := range tests {
		actual, err := Marshal(test.c)
		if err != nil {
			t.Fatalf("return err %v", err.Error())
		}

		if string(actual) != test.want {
			t.Errorf("Result = %s, Expected %v", string(actual), test.c)
		}
	}
}

func TestComplex128(t *testing.T) {
	var tests = []struct {
		c    complex128
		want string
	}{
		{1 - 2i, "#C(1.000000 -2.000000)"},
	}
	for _, test := range tests {
		actual, err := Marshal(test.c)
		if err != nil {
			t.Fatalf("return err %v", err.Error())
		}

		if string(actual) != test.want {
			t.Errorf("Result = %s, Expected %v", string(actual), test.c)
		}
	}
}

func TestBool(t *testing.T) {
	var tests = []struct {
		b    bool
		want string
	}{
		{true, "t"},
		{false, "nil"},
	}
	for _, test := range tests {
		actual, err := Marshal(test.b)
		if err != nil {
			t.Fatalf("return err %v", err.Error())
		}
		if string(actual) != test.want {
			t.Errorf("Result = %s, Expected %v", actual, test.want)
		}
	}
}
