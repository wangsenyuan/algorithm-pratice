package lcp79

var dd = []int{-1, 0, 1, 0, -1}

func extractMantra(matrix []string, mantra string) int {
	m := len(matrix)
	n := len(matrix[0])
	ln := len(mantra) + 1
	// dp[i][j][k] 表示 hand 位于位置(i,j)，且符文当前位置在k时的最小操作次数
	dp := make([][][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([][]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = make([]int, ln)
			for k := 0; k < ln; k++ {
				dp[i][j][k] = -1
			}
		}
	}

	que := make([]int, m*n*ln)
	var front, end int
	pushBack := func(i, j, k int) {
		cur := i*n*ln + j*ln + k
		que[end] = cur
		end++
	}

	popFront := func() (i int, j int, k int) {
		cur := que[front]
		front++
		i = cur / (n * ln)
		j = (cur % (n * ln)) / ln
		k = cur % ln
		return
	}

	dp[0][0][0] = 0
	pushBack(0, 0, 0)
	//
	//if matrix[0][0] == mantra[0] {
	//	dp[0][0][1] = 1
	//	pushBack(0, 0, 1)
	//}

	for front < end {
		r, c, p := popFront()
		if p < len(mantra) && matrix[r][c] == mantra[p] && dp[r][c][p+1] < 0 {
			dp[r][c][p+1] = dp[r][c][p] + 1
			pushBack(r, c, p+1)
		}
		for k := 0; k < 4; k++ {
			u, v := r+dd[k], c+dd[k+1]
			if u < 0 || u == m || v < 0 || v == n {
				continue
			}
			// not pick
			if dp[u][v][p] < 0 {
				dp[u][v][p] = dp[r][c][p] + 1
				pushBack(u, v, p)
			}
		}
	}
	res := -1

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if dp[i][j][ln-1] > 0 && (res < 0 || res > dp[i][j][ln-1]) {
				res = dp[i][j][ln-1]
			}
		}
	}
	return res
}
