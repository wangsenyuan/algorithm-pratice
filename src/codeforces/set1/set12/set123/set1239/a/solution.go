package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	res := solve(n, m)

	fmt.Println(res)
}

const mod = 1e9 + 7

func add(a, b int) int {
	a += b
	if a >= mod {
		a -= mod
	}
	return a
}

func sub(a, b int) int {
	return add(a, mod-b)
}

func mul(a, b int) int {
	return a * b % mod
}

func solve(n int, m int) int {
	if n < m {
		n, m = m, n
	}
	dp := make([]int, n+3)
	dp[0] = 1
	dp[1] = 1
	for i := 2; i <= n; i++ {
		dp[i] = add(dp[i-2], dp[i-1])
	}

	res := sub(add(dp[n], dp[m]), 1)
	res = mul(res, 2)

	return res
}
