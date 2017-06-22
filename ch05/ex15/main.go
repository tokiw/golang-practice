package main

import "fmt"

//!+
func sum(vals ...int) int {
	total := 0
	for _, val := range vals {
		total += val
	}
	return total
}

func max(vals ...int) int {
	if len(vals) == 0 {
		panic("no argument")
	}
	max := vals[0]
	for _, val := range vals {
		if max < val {
			max = val
		}
	}
	return max
}

func max2(v int, vals ...int) int {
	for _, val := range vals {
		if v < val {
			v = val
		}
	}
	return v
}

func min(vals ...int) int {
	if len(vals) == 0 {
		panic("no argument")
	}
	min := vals[0]
	for _, val := range vals {
		if min > val {
			min = val
		}
	}
	return min
}

func min2(v int, vals ...int) int {
	for _, val := range vals {
		if v > val {
			v = val
		}
	}
	return v
}

//!-

func main() {
	//!+main
	fmt.Println(sum())           //  "0"
	fmt.Println(sum(3))          //  "3"
	fmt.Println(sum(1, 2, 3, 4)) //  "10"
	//!-main

	//!+slice
	values := []int{1, 2, 3, 4}
	fmt.Println(sum(values...)) // "10"
	//!-slice
}
