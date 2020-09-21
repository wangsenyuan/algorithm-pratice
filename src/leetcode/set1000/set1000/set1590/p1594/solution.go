package p1594

func maxUniqueSplit(s string) int {
	n := len(s)
	// n <= 16
	m := n - 1
	M := 1 << m
	// a|b|c*d|d
	var best int
	for state := 1; state < M; state++ {
		mem := make(map[string]bool)
		var j = -1
		for i := 0; i < m; i++ {
			if state&(1<<i) > 0 {
				x := m - i
				y := m - j

				tmp := s[x:y]

				mem[tmp] = true

				j = i
			}
		}
		if j < m {
			tmp := s[:m-j]
			mem[tmp] = true
		}
		best = max(best, len(mem))
	}

	return best
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
