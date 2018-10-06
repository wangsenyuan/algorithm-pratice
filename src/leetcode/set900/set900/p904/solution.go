package p904

func totalFruit(tree []int) int {
	wnd := make(map[int]int)

	n := len(tree)

	var res int
	for i, j := 0, 0; i < n; i++ {
		wnd[tree[i]]++
		for len(wnd) > 2 {
			wnd[tree[j]]--
			if wnd[tree[j]] == 0 {
				delete(wnd, tree[j])
			}
			j++
		}
		res = max(res, i-j+1)
	}

	return res
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
