package p1952

func numberOfWeeks(milestones []int) int64 {
	n := len(milestones)
	var x int
	var sum int64
	for i := 0; i < n; i++ {
		x = max(x, milestones[i])
		sum += int64(milestones[i])
	}
	y := sum - int64(x)
	return min(2*y+1, sum)
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func min(a, b int64) int64 {
	if a <= b {
		return a
	}
	return b
}
