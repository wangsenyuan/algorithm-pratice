package p849

func maxDistToClosest(seats []int) int {
	n := len(seats)
	left := make([]int, n)
	right := make([]int, n)

	for i := 0; i < n; i++ {
		left[i] = i
		if seats[i] == 0 {
			if i > 0 {
				left[i] = left[i-1]
			} else {
				left[i] = -n
			}
		}
	}

	for i := n - 1; i >= 0; i-- {
		right[i] = i
		if seats[i] == 0 {
			if i < n-1 {
				right[i] = right[i+1]
			} else {
				right[i] = 2 * n
			}
		}
	}

	var best int

	for i := 0; i < n; i++ {
		j := left[i]
		k := right[i]
		a := i - j
		b := k - i
		c := min(a, b)
		best = max(c, best)
	}

	return best
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
