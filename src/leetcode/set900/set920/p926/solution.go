package p926

func minFlipsMonoIncr(S string) int {
	n := len(S)
	f := make([]int, n)
	var cnt int
	for i := 0; i < n; i++ {
		if S[i] == '1' {
			cnt++
		}
		f[i] = cnt
	}

	g := make([]int, n)
	cnt = 0
	for i := n - 1; i >= 0; i-- {
		if S[i] == '1' {
			cnt++
		}
		g[i] = cnt
	}
	// make 1 at pos 0, then need to flip n - g[0]
	best := n - g[0]
	for i := 1; i < n; i++ {
		//flip S[i] to 1
		best = min(best, f[i-1]+n-i-g[i])
	}
	best = min(best, f[n-1])
	return best
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
