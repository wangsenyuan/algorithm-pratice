package p1376

func numOfMinutes(n int, headID int, manager []int, informTime []int) int {
	conns := make([][]int, n)

	for i := 0; i < n; i++ {
		conns[i] = make([]int, 0, 10)
	}

	for i := 0; i < n; i++ {
		j := manager[i]
		if j >= 0 {
			conns[j] = append(conns[j], i)
		}
	}

	var dfs func(u int, cur int) int

	dfs = func(u int, cur int) int {
		if len(conns[u]) == 0 {
			return cur
		}

		var res int

		for _, v := range conns[u] {
			x := dfs(v, cur+informTime[u])
			if x > res {
				res = x
			}
		}
		return res
	}

	return dfs(headID, 0)
}
