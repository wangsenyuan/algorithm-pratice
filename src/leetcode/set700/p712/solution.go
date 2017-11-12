package main

import "fmt"

func minimumDeleteSum(s1 string, s2 string) int {
	n := len(s1)
	m := len(s2)

	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, m+1)
	}
	dp[0][0] = 0
	for j := 1; j <= m; j++ {
		dp[0][j] = int(s2[j-1]) + dp[0][j-1]
	}

	for i := 1; i <= n; i++ {
		dp[i][0] = dp[i-1][0] + int(s1[i-1])
		for j := 1; j <= m; j++ {
			if s1[i-1] == s2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				a := dp[i-1][j] + int(s1[i-1])
				b := dp[i][j-1] + int(s2[j-1])
				dp[i][j] = min(a, b)
			}
		}
	}

	return dp[n][m]
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func main() {
	//println(int('a'))
	fmt.Println(minimumDeleteSum("sea", "eat"))
	fmt.Println(minimumDeleteSum("delete", "leet"))
}
