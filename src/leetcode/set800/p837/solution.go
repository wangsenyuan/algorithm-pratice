package p837

func new21Game(N int, K int, W int) float64 {
	f := make([]float64, N+3)
	f[0] = 1.0
	f[1] = -1

	var val float64
	for i := 0; i < K; i++ {
		val += f[i]
		if i+1 <= N {
			f[i+1] += val / float64(W)
		}
		if i+W+1 <= N {
			f[i+W+1] -= val / float64(W)
		}
	}
	var res float64
	for i := K; i <= N; i++ {
		val += f[i]
		res += val
	}
	return res
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
