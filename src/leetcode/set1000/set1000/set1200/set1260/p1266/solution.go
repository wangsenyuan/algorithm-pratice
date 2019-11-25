package p1266

func minTimeToVisitAllPoints(points [][]int) int {
	var res int

	for i := 1; i < len(points); i++ {
		res += max(abs(points[i][0]-points[i-1][0]), abs(points[i][1]-points[i-1][1]))
	}

	return res
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
