package p879

const MOD = 1000000007

func profitableSchemes(G int, P int, group []int, profit []int) int {
	n := len(group)

	createDpTable := func() [][]int64 {
		dp := make([][]int64, G+1)
		for i := 0; i <= G; i++ {
			dp[i] = make([]int64, P+1)
		}
		return dp
	}

	add := func(a *int64, b int64) {
		*a += b
		if *a >= MOD {
			*a -= MOD
		}
	}

	dp := createDpTable()
	dp[0][0] = 1
	for i := 0; i < n; i++ {
		ndp := createDpTable()
		for j := G; j >= 0; j-- {
			for k := P; k >= 0; k-- {
				add(&ndp[j][k], dp[j][k])
				nk := min(P, k+profit[i])
				nj := j + group[i]
				if nj <= G {
					add(&ndp[nj][nk], dp[j][k])
				}
			}
		}
		dp = ndp
	}

	var ret int64

	for j := 0; j <= G; j++ {
		add(&ret, dp[j][P])
	}

	return int(ret)
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
