package p2850

func isReachableAtTime(sx int, sy int, fx int, fy int, t int) bool {
	if sx == fx && sy == fy {
		return t != 1
	}

	dx := abs(sx - fx)
	dy := abs(sy - fy)
	min_time := max(dx, dy)
	return min_time <= t
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
