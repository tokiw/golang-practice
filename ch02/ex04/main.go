package popcount

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	var count int
	for i := 0; i < 64; i++ {
		tmp := x >> uint(i)
		if tmp&1 == 1 {
			count++
		}
	}
	return count
}
