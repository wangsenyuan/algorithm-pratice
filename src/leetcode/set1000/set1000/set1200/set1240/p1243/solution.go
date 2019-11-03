package p1243

func transformArray(arr []int) []int {
	n := len(arr)

	if n <= 2 {
		return arr
	}

	res := make([]int, n)
	copy(res, arr)

	bak := make([]int, n)

	for {
		copy(bak, res)
		var change bool
		for i := 1; i < n-1; i++ {
			if bak[i-1] > bak[i] && bak[i+1] > bak[i] {
				change = true
				res[i]++
			} else if bak[i-1] < bak[i] && bak[i+1] < bak[i] {
				change = true
				res[i]--
			}
		}
		if !change {
			return res
		}
	}
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
