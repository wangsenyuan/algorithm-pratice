package p1030

func longestArithSeqLength(A []int) int {
	n := len(A)
	if n <= 2 {
		return n
	}
	f := make([]map[int]int, n)
	for i := 0; i < n; i++ {
		f[i] = make(map[int]int)

		for j := 0; j < i; j++ {
			f[i][A[i]-A[j]] = 2
		}
	}

	var best = 2
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			d := A[i] - A[j]
			if x, found := f[j][d]; found {
				f[i][d] = max(f[i][d], x+1)
			}
			best = max(best, f[i][d])
		}
	}

	return best
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
