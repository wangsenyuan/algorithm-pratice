package main

import "fmt"

func numDistinct(s string, t string) int {
	m := len(s)
	n := len(t)

	if m < n {
		return 0
	}

	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	for i, x := range s {
		for j, y := range t {
			if x != y {
				continue
			}

			if j == 0 {
				dp[i][j] = 1
				continue
			}
			for k := j - 1; k < i; k++ {
				dp[i][j] += dp[k][j - 1]
			}
		}
	}

	rs := 0

	for i := n - 1; i < m; i++ {
		rs += dp[i][n - 1]
	}

	return rs
}

func main() {
	fmt.Println(numDistinct("ddd", "dd"))
}
