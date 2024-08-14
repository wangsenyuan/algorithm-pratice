package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(solve(n))
}

const N = 1_000_000 + 10
const mod = 998244353

func add(a, b int) int {
	a += b
	if a >= mod {
		a -= mod
	}
	return a
}

func solve(n int) int {
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		for j := 2 * i; j <= n; j += i {
			dp[j] = add(dp[j], 1)
		}
	}

	dp[0] = 1
	sum := 1
	for i := 1; i <= n; i++ {
		dp[i] = add(dp[i], sum)
		sum = add(sum, dp[i])
	}

	return dp[n]
}
