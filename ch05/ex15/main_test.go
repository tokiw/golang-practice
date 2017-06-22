package main

import (
	"testing"
)

func TestMax(t *testing.T) {
	assertMax(t, 4, 1, 2, 3, 4)
	assertMax(t, 4, 4, 4, 4, 4)
	assertMax(t, 4, 4, 3, 2, 1)
	assertMax(t, -1, -4, -3, -2, -1)
	assertMax(t, -1, -1, -2, -3, -4)
	assertMaxPanic(t)
}

func TestMax2(t *testing.T) {
	assertMax2(t, 4, 1, 2, 3, 4)
	assertMax2(t, 4, 4, 4, 4, 4)
	assertMax2(t, 4, 4, 3, 2, 1)
	assertMax2(t, -1, -4, -3, -2, -1)
	assertMax2(t, -1, -1, -2, -3, -4)
}

func TestMin(t *testing.T) {
	assertMin(t, 1, 1, 2, 3, 4)
	assertMin(t, 4, 4, 4, 4, 4)
	assertMin(t, 1, 4, 3, 2, 1)
	assertMin(t, -4, -4, -3, -2, -1)
	assertMin(t, -4, -1, -2, -3, -4)
	assertMinPanic(t)
}

func TestMin2(t *testing.T) {
	assertMin2(t, 1, 1, 2, 3, 4)
	assertMin2(t, 4, 4, 4, 4, 4)
	assertMin2(t, 1, 4, 3, 2, 1)
	assertMin2(t, -4, -4, -3, -2, -1)
	assertMin2(t, -4, -1, -2, -3, -4)
}

func assertMax(t *testing.T, expected int, vals ...int) {
	if r := max(vals...); r != expected {
		t.Errorf("result: %v, expected: %v", r, expected)
	}
}

func assertMaxPanic(t *testing.T) {
	defer func() {
		if p := recover(); p != nil {
			if p != "no argument" {
				t.Errorf("result: %v", p)
			}
		} else {
			t.Errorf("no panic")
		}
	}()
	max()
}

func assertMin(t *testing.T, expected int, vals ...int) {
	if r := min(vals...); r != expected {
		t.Errorf("result: %v, expected: %v", r, expected)
	}
}

func assertMinPanic(t *testing.T) {
	defer func() {
		if p := recover(); p != nil {
			if p != "no argument" {
				t.Errorf("result: %v", p)
			}
		} else {
			t.Errorf("no panic")
		}
	}()
	min()
}

func assertMax2(t *testing.T, expected int, val int, vals ...int) {
	if r := max2(val, vals...); r != expected {
		t.Errorf("result: %v, expected: %v", r, expected)
	}
}

func assertMin2(t *testing.T, expected int, val int, vals ...int) {
	if r := min2(val, vals...); r != expected {
		t.Errorf("result: %v, expected: %v", r, expected)
	}
}
