package main

import (
	"fmt"
	"sync"
)

// pc[i] is the population count of i.
var pc [256]byte
var initTableOnce sync.Once

func initTable() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	initTableOnce.Do(initTable)

	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func main() {
	for i := 0; i < 10; i++ {
		x := i ^ (1 << uint(i))
		fmt.Printf("x = %d, popcount = %d\n", x, PopCount(uint64(x)))
	}
}
