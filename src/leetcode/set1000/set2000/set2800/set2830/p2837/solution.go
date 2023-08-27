package p2837

func furthestDistanceFromOrigin(moves string) int {
	var l int
	var r int
	n := len(moves)

	for i := 0; i < n; i++ {
		if moves[i] == 'L' {
			l++
		} else if moves[i] == 'R' {
			r++
		}
	}
	return abs(l-r) + n - (l + r)
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
