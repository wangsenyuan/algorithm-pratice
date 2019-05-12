package p1042

func gardenNoAdj(N int, paths [][]int) []int {
	outs := make([][]int, N)
	for i := 0; i < N; i++ {
		outs[i] = make([]int, 0, 3)
	}

	for _, path := range paths {
		u, v := path[0] - 1, path[1] - 1
		outs[u] = append(outs[u], v)
		outs[v] = append(outs[v], u)
	}

	ans := make([]int, N)

	fail := make([]int, N)

	var dfs func(u int) int

	dfs = func(u int) int {
		if ans[u] != 0 {
			return ans[u]
		}

		for color := 1; color <= 4; color++ {
			ans[u] = color

			if fail[u]&(1<<uint(color)) > 0 {
				continue
			}

			can := true
			for _, v := range outs[u] {
				tmp := dfs(v)
				if tmp == color {
					can = false
					break
				}
			}
			if can {
				return color
			}
			fail[u] |= 1 << uint(color)
		}

		return 0
	}

	for i := 0; i < N; i++ {
		if ans[i] == 0 {
			dfs(i)
		}
	}

	return ans
}
