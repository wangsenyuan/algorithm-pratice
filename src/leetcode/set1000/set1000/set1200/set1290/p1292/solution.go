package p1292

func maxSideLength(mat [][]int, threshold int) int {
	m := len(mat)
	n := len(mat[0])
	sum := make([][]int, m)
	for i := 0; i < m; i++ {
		sum[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			sum[i][j] = mat[i][j]
			if i > 0 {
				sum[i][j] += sum[i-1][j]
			}
			if j > 0 {
				sum[i][j] += sum[i][j-1]
			}
			if i > 0 && j > 0 {
				sum[i][j] -= sum[i-1][j-1]
			}
		}
	}
	var best int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			for k := 0; i+k < m && j+k < n; k++ {
				tmp := sum[i+k][j+k]
				if i > 0 {
					tmp -= sum[i-1][j+k]
				}
				if j > 0 {
					tmp -= sum[i+k][j-1]
				}
				if i > 0 && j > 0 {
					tmp += sum[i-1][j-1]
				}
				if tmp > threshold {
					break
				}
				best = max(best, k+1)
			}
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
