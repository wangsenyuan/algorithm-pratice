package p754

import "math"

func reachNumber(target int) int {
	target = abs(target)

	n := int64((-1.0 + math.Sqrt(1.0+8.0*float64(target))) / 2.0)
	sum := n * (n + 1) / 2
	if sum < int64(target) {
		n++
	}
	sum = n * (n + 1) / 2
	if sum == int64(target) {
		return int(n)
	}

	left := sum - int64(target)
	if left&1 == 0 {
		return int(n)
	}

	if n&1 == 1 {
		return int(n + 2)
	}
	return int(n + 1)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
