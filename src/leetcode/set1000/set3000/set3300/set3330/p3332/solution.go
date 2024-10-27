package p3332

import "slices"

func maxScore(n int, k int, stayScore [][]int, travelScore [][]int) int {
	dp := make([]int, n)
	fp := make([]int, n)

	for i := 0; i < k; i++ {
		for j := 0; j < n; j++ {
			// 它前一天待在城市j, 且今天仍然在城市j
			fp[j] = max(fp[j], dp[j]+stayScore[i][j])
			for k := 0; k < n; k++ {
				if j != k {
					fp[k] = max(fp[k], dp[j]+travelScore[j][k])
				}
			}
		}
		copy(dp, fp)
		clear(fp)
	}
	return slices.Max(dp)
}
