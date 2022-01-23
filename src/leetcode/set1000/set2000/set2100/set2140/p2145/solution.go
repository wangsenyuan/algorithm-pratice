package p2145

func numberOfArrays(differences []int, lower int, upper int) int {
	n := len(differences)

	var a, b int
	var cur int
	for i := 1; i <= n; i++ {
		cur += differences[i-1]
		a = max(a, cur)
		b = min(b, cur)
	}

	if a-b > upper-lower {
		return 0
	}
	return (upper - lower) - (a - b) + 1
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
