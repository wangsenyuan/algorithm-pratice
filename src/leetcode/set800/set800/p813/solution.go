package p813

func largestSumOfAverages(A []int, K int) float64 {
	n := len(A)

	f := make([][]float64, n)

	for i := 0; i < n; i++ {
		f[i] = make([]float64, K+1)
		for j := 0; j <= K; j++ {
			f[i][j] = -1
		}
	}

	for k := 1; k <= K; k++ {
		for i := 0; i < n; i++ {
			var sum int
			for j := i; j >= 0; j-- {
				sum += A[j]
				if k == 1 && j == 0 {
					f[i][k] = float64(sum) / float64(i+1)
					continue
				}
				if j == 0 || f[j-1][k-1] < 0 {
					continue
				}
				if f[i][k] < 0 || f[i][k] < f[j-1][k-1]+float64(sum)/float64(i-j+1) {
					f[i][k] = f[j-1][k-1] + float64(sum)/float64(i-j+1)
				}
			}
		}
	}

	return f[n-1][K]
}
