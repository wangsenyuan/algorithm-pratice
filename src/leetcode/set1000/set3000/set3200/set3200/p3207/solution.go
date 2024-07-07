package p3207

func numberOfAlternatingGroups(colors []int, k int) int {
	n := len(colors)
	dc := make([]int, n*2)
	copy(dc, colors)
	copy(dc[n:], colors)

	p2 := make([]int, 2*n)
	vis := make([]bool, n)
	p2[0] = -1
	var res int
	for i := 1; i < 2*n; i++ {
		if dc[i] == dc[i-1] {
			p2[i] = i - 1
		} else if i == 1 || dc[i] != dc[i-2] {
			p2[i] = i - 2
		} else {
			p2[i] = max(p2[i-2], p2[i-1])
		}
		if i-p2[i] >= k && !vis[i%n] {
			res++
			vis[i%n] = true
		}
	}

	return res
}
