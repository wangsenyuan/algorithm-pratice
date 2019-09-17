package p1192

func criticalConnections(n int, connections [][]int) [][]int {
	outs := make([][]int, n)

	for i := 0; i < n; i++ {
		outs[i] = make([]int, 0, 10)
	}

	for _, connection := range connections {
		a, b := connection[0], connection[1]
		outs[a] = append(outs[a], b)
		outs[b] = append(outs[b], a)
	}

	euler := make([]int, n)
	low := make([]int, n)
	for i := 0; i < n; i++ {
		euler[i] = -1
		low[i] = -1
	}

	var dfs func(p int, u int, time *int)

	var res [][]int

	dfs = func(p int, u int, time *int) {
		if euler[u] >= 0 {
			return
		}
		euler[u] = *time
		*time++
		low[u] = euler[u]

		for _, v := range outs[u] {
			if p == v {
				continue
			}
			dfs(u, v, time)
			if low[v] > euler[u] {
				res = append(res, []int{u, v})
			}
			low[u] = min(low[u], low[v])
		}
	}

	// dfs = func(p int, u int, time *int) int {
	// 	euler[u] = *time
	// 	*time++

	// 	rt := n

	// 	for _, v := range outs[u] {
	// 		if p == v {
	// 			continue
	// 		}
	// 		if euler[v] >= 0 {
	// 			rt = min(rt, euler[v])
	// 		} else {
	// 			tmp := dfs(u, v, time)
	// 			if tmp > euler[u] {
	// 				res = append(res, []int{u, v})
	// 			}
	// 			rt = min(rt, tmp)
	// 		}
	// 	}

	// 	return rt
	// }

	dfs(-1, 0, new(int))

	return res
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
