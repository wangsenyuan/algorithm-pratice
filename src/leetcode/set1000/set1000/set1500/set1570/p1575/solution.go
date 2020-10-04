package p1575

const MOD = 1000000007

func countRoutes(locations []int, start int, finish int, fuel int) int {
	n := len(locations)

	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, fuel+1)
	}

	dp[start][fuel] = 1

	for k := fuel; k > 0; k-- {
		for i := 0; i < n; i++ {
			if dp[i][k] == 0 {
				continue
			}
			for j := 0; j < n; j++ {
				d := abs(locations[i] - locations[j])
				if i == j || d > k {
					continue
				}
				dp[j][k-d] += dp[i][k]
				if dp[j][k-d] >= MOD {
					dp[j][k-d] -= MOD
				}
			}
		}
	}

	var res int

	for k := fuel; k >= 0; k-- {
		res += dp[finish][k]
		if res >= MOD {
			res -= MOD
		}
	}
	return res
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
