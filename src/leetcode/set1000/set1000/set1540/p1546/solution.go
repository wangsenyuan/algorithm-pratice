package p1546

import "sort"

func minCost(n int, cuts []int) int {
	sort.Ints(cuts)

	leaf := make([]int, 0, len(cuts)+1)

	for i := 0; i < len(cuts); i++ {
		if i == 0 {
			leaf = append(leaf, cuts[i])
		} else {
			leaf = append(leaf, cuts[i]-cuts[i-1])
		}
	}
	leaf = append(leaf, n-cuts[len(cuts)-1])

	//dp[i][j] = min_cost join leaf[i] to leaf[j]
	m := len(leaf)

	sums := make([]int, m+1)
	for i := 1; i <= m; i++ {
		sums[i] = sums[i-1] + leaf[i-1]
	}

	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, m)
		for j := 0; j < m; j++ {
			dp[i][j] = -1
		}
		dp[i][i] = 0
	}

	for k := 2; k <= m; k++ {
		for i := 0; i+k <= m; i++ {
			j := i + k - 1
			for kk := 1; kk < k; kk++ {
				tmp := dp[i][i+kk-1] + dp[i+kk][j] + sums[j+1] - sums[i]
				if dp[i][j] < 0 || tmp < dp[i][j] {
					dp[i][j] = tmp
				}
			}
		}
	}

	return dp[0][m-1]
}
