package main

import (
	"crypto/sha256"
	"fmt"

	"github.com/tokiw/golang-practice/ch02/ex03"
)

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Printf("%x\n%x\ndiff: %v", c1, c2, bitDiff(c1, c2))
}

func bitDiff(b1, b2 [32]byte) int {
	var count int

	for i := 0; i < 32; i++ {
		count += popcount.PopCount(uint64(b1[i] ^ b2[i]))
	}
	return count
}
