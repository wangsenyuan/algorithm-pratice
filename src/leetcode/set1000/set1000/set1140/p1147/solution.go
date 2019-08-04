package p1147

func longestDecomposition(text string) int {
	n := len(text)
	mem := make(map[int]int)

	var dfs func(i int, j int) int

	dfs = func(i int, j int) int {
		if i >= j {
			return 0
		}

		key := i*n + j
		if ans, found := mem[key]; found {
			return ans
		}

		best := 1

		for u := i + 1; u < j; u++ {
			l := u - i

			if j-l < u {
				break
			}
			pref := text[i:u]
			v, a := j-l, 0
			for v < j && a < l && pref[a] == text[v] {
				v++
				a++
			}

			if l == a {
				// a candiate
				best = max(best, dfs(u, j-l)+2)
			}
		}

		mem[key] = best
		return best
	}

	return dfs(0, n)
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
