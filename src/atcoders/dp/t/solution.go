package main

import "fmt"

func main() {

	var n int
	fmt.Scan(&n)
	var s string
	fmt.Scan(&s)

	res := solve(s)

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

func mul(a, b int) int {
	return a * b % mod
}

func solve(s string) int {
	n := len(s) + 1
	dp := make([]int, n+1)
	dp[1] = 1

	for i := 1; i < n; i++ {
		if s[i-1] == '<' {
			var sum int
			for j := 1; j <= i; j++ {
				tmp := add(sum, dp[j])
				dp[j] = sum
				sum = tmp
			}
			dp[i+1] = sum
		} else {
			var sum int
			for j := i; j > 0; j-- {
				dp[j] = add(sum, dp[j])
				sum = dp[j]
			}
		}
	}

	var res int
	for j := 1; j <= n; j++ {
		res = add(res, dp[j])
	}
	return res
}
