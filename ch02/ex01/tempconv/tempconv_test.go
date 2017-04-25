package tempconv

import "testing"

func TestCToF(t *testing.T) {
	actual := CToF(FreezingC)
	expected := Fahrenheit(32)
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestCToK(t *testing.T) {
	actual := CToK(AbsoluteZeroC)
	expected := Kelvin(0)
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestFToC(t *testing.T) {
	actual := FToC(32)
	expected := FreezingC
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestFToK(t *testing.T) {
	actual := FToK(32)
	expected := Kelvin(273.15)
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestKToC(t *testing.T) {
	actual := KToC(0)
	expected := AbsoluteZeroC
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestKToF(t *testing.T) {
	actual := KToF(273.15)
	expected := Fahrenheit(32)
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}
