package p808

func soupServings(N int) float64 {
	if N >= 5000 {
		return 1.0
	}

	var f func(a, b int) float64
	n := (N + 24) / 25
	cache := make([][]float64, n+1)
	for i := 0; i <= n; i++ {
		cache[i] = make([]float64, n+1)
		for j := 0; j <= n; j++ {
			cache[i][j] = -1
		}
	}

	f = func(a, b int) float64 {
		if a <= 0 && b <= 0 {
			return 0.5
		}

		if a <= 0 {
			return 1.0
		}

		if b <= 0 {
			return 0.0
		}

		if cache[a][b] >= 0 {
			return cache[a][b]
		}

		cache[a][b] = 0.25 * (f(a-4, b) + f(a-3, b-1) + f(a-2, b-2) + f(a-1, b-3))

		return cache[a][b]
	}

	return f(n, n)
}
