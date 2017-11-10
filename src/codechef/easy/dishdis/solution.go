package main

import "fmt"

const MOD = 1000000007

func main() {
	var t int
	fmt.Scanf("%d", &t)

	for i := 0; i < t; i++ {
		solve()
	}
}

func solve() {
	var n, m int
	fmt.Scanf("%d %d", &n, &m)
	rs := make([][]int, m)

	for i := 0; i < m; i++ {
		rs[i] = make([]int, 2)
		fmt.Scanf("%d %d", &rs[i][0], &rs[i][1])
	}

	dp := make([][]int, m+1)

	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}

	dp[0][0] = 1
	for i := 1; i <= m; i++ {
		a, b := rs[i-1][0], rs[i-1][1]
		for j := a; j <= n; j++ {
			tmp := 0
			for k := a; k <= b && k <= j; k++ {
				tmp += dp[i-1][j-k]
				if tmp >= MOD {
					tmp -= MOD
				}
			}
			dp[i][j] = tmp
		}
	}
	ans := dp[m][n]

	fmt.Println(ans)
}
