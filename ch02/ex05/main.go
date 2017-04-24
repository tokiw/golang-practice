package popcount

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	var count int
	for x > 0 {
		tmp := x & (x - 1)
		if x > tmp {
			count++
		}
		x = tmp
	}
	return count
}
