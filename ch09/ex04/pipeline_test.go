package pipeline

import (
	"testing"
)

func BenchmarkPipeline10(b *testing.B) {
	in, out := pipeline(10)
	for i := 0; i < b.N; i++ {
		in <- 1
		<-out
	}
	close(in)
}

func BenchmarkPipeline10000(b *testing.B) {
	in, out := pipeline(10000)
	for i := 0; i < b.N; i++ {
		in <- 1
		<-out
	}
	close(in)
}

func BenchmarkPipeline100000(b *testing.B) {
	in, out := pipeline(1000000)
	for i := 0; i < b.N; i++ {
		in <- 1
		<-out
	}
	close(in)
}

// func BenchmarkPipeline1000000(b *testing.B) {
// 	in, out := pipeline(10000000)
// 	for i := 0; i < b.N; i++ {
// 		in <- 1
// 		<-out
// 	}
// 	close(in)
// }
