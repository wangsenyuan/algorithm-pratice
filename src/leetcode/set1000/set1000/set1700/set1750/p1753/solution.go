package p1753

import "sort"

func maximumScore(a int, b int, c int) int {
	arr := []int{a, b, c}
	sort.Ints(arr)
	a, b, c = arr[0], arr[1], arr[2]
	// x + y <= a
	// y + z <= b
	// x + z <= c
	// y + min(x, z - y)
	var best int
	for y := 1; y <= b; y++ {
		x := min(a, c-y)
		// x of a pair with c
		tmp := x + y
		if x < a {
			tmp += min(b-y, a-x)
		}
		best = max(best, tmp)
	}

	return best
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
