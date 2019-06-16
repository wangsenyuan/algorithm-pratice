package main

import (
	"fmt"
)

func findLength(A []int, B []int) int {
	n := len(A)
	m := len(B)

	dp := make([][]int, n+1)

	for i := 0; i <= n; i++ {
		dp[i] = make([]int, m+1)
	}

	ans := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if A[i] == B[j] {
				dp[i+1][j+1] = 1 + dp[i][j]
				if dp[i+1][j+1] > ans {
					ans = dp[i+1][j+1]
				}
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(findLength([]int{1, 2, 3, 2, 1}, []int{3, 2, 1, 4, 7}))
}
