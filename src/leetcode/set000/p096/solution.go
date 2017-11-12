package main

import "fmt"

func numTrees(n int) int {
	dp := make([][]int, n)

	for i := range dp {
		dp[i] = make([]int, n + 1)
		dp[i][1] = 1
	}
	for k := 2; k <= n; k++ {
		for i := 0; i <= n - k; i++ {
			for l := 0; l < k; l++ {
				if l == 0 {
					dp[i][k] += dp[i + 1][k - 1]
				} else if l == k - 1 {
					dp[i][k] += dp[i][k - 1]
				} else if i + l + 1 < n {
					dp[i][k] += dp[i][l] * dp[i + l + 1][k - l - 1]
				}
			}
		}
	}
	return dp[0][n]
}

func main() {
	fmt.Println(numTrees(5))
}
