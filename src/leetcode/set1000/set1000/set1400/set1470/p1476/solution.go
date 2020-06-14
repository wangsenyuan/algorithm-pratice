package p1476

import "sort"

const INF = 1 << 30

func minDistance(H []int, K int) int {
	sort.Ints(H)

	n := len(H)

	S := make([]int, n)
	for i := 0; i < n; i++ {
		S[i] = H[i]
		if i > 0 {
			S[i] += S[i-1]
		}
	}

	DD := make([][]int, n)
	for i := 0; i < n; i++ {
		DD[i] = make([]int, n)
		for j := i + 1; j < n; j++ {
			var x int
			ii := i
			for ii < j && H[ii]-H[i] < H[j]-H[ii] {
				x += H[ii] - H[i]
				ii++
			}
			var y int
			for ii < j {
				y += H[j] - H[ii]
				ii++
			}
			DD[i][j] = x + y
		}
	}

	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, K+1)
		for kk := 0; kk <= K; kk++ {
			dp[i][kk] = INF
		}
	}

	// fp := make([][]int, n)
	// for i := 0; i < n; i++ {
	// 	fp[i] = make([]int, K+1)
	// 	for k := 0; k <= K; k++ {
	// 		fp[i][k] = INF
	// 	}
	// }

	dp[0][1] = 0
	for i := 1; i < n; i++ {
		// if we set a station at house i
		for k := 1; k <= i && k < K; k++ {
			for j := i - 1; j >= 0; j-- {
				dp[i][k+1] = min(dp[i][k+1], dp[j][k]+DD[j][i])
			}
		}
		dp[i][1] = (i+1)*H[i] - S[i]
	}

	var res = dp[n-1][K]

	var x int
	for j := n - 1; j >= 0; j-- {
		x += H[j]
		tmp := dp[j][K] + x - (n-1-j+1)*H[j]
		res = min(res, tmp)
	}

	return res
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
