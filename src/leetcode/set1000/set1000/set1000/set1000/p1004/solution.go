package p1004

func longestOnes(A []int, K int) int {
	var best int
	var j int
	n := len(A)
	for i := 0; i < n; i++ {
		if A[i] == 1 || K > 0 {
			best = max(best, i-j+1)
			K -= 1 - A[i]
			continue
		}
		// K == 0
		for A[j] == 1 {
			j++
		}
		// A[j] == 0
		j++
		best = max(best, i-j+1)
	}

	return best
}

func longestOnes1(A []int, K int) int {
	n := len(A)

	f := make([]int, n)
	g := make([]int, n)
	if A[n-1] == 0 {
		f[n-1] = n
		g[n-1] = n - 1
	} else {
		f[n-1] = n - 1
		g[n-1] = n
	}

	for i := n - 2; i >= 0; i-- {
		if A[i] == 0 {
			f[i] = f[i+1]
			g[i] = i
		} else {
			f[i] = i
			g[i] = g[i+1]
		}
	}

	flip := func(i int) int {
		// i == n || A[i] == 0
		k := K
		j := i
		for j < n && j+k >= f[j] {
			k -= f[j] - j
			j = f[j]
			if j < n && A[j] == 1 {
				// skip ones
				j = g[j]
			}
		}

		if j+k <= n {
			return j + k - i
		}

		return j - i
	}

	var best int

	for i := 0; i < n; {
		//flip K zerons from position i
		var tmp int
		if A[i] == 0 {
			tmp = flip(i)
			i++
		} else {
			tmp = g[i] - i + flip(g[i])
			i = g[i]
		}
		best = max(best, tmp)
	}

	return best
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
