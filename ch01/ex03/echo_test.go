package main

import "testing"

var slice = []string{"a", "b", "c", "d", "e", "f"}

func BenchmarkIsEcho(b *testing.B) {
	for i := 0; i < b.N; i++ {
		echo(slice)
	}
}

func BenchmarkIsEcho1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		echo1(slice)
	}
}
