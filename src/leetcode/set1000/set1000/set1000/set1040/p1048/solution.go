package p1048

func longestStrChain(words []string) int {
	n := len(words)
	succ := make([][]int, n)
	for i := 0; i < n; i++ {
		succ[i] = make([]int, 0, 10)
 		for j := 0; j < n; j++ {
			if predecessor(words[i], words[j]) {
				succ[i] = append(succ[i], j)
			}
		}
	}

	var dfs func(u int) int

	mem := make([]int, n)

	dfs = func(u int) int {
		if mem[u] > 0 {
			return mem[u]
		}
		best := 1

		for _, v := range succ[u] {
			best = max(best, dfs(v) + 1)
		}
		mem[u] = best
		return best
	}

	var res int

	for i := 0; i < n; i++ {
		res = max(res, dfs(i))
	}

	return res
}

func max(a int, b int) int {
	if a >= b {
		return a
	}
	return b
}


func predecessor(a, b string) bool {
	if len(a) + 1 != len(b) {
		return false
	}

	var i int

	for i < len(a) && a[i] == b[i] {
		i++
	}

	for i < len(a) && a[i] == b[i+1] {
		i++
	}

	return i == len(a)
}