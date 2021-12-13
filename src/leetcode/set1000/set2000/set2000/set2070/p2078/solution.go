package p2078

func maxDistance(colors []int) int {
	var res int
	n := len(colors)
	for i := 0; i < n; i++ {
		for j := n - 1; j > i; j-- {
			if colors[i] != colors[j] {
				res = max(res, j-i)
				break
			}
		}
	}

	return res
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
