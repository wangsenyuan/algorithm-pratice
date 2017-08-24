package main

import "fmt"

func strangePrinter(s string) int {
	n := len(s)

	if n == 0 {
		return 0
	}

	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
		dp[i][i] = 1
	}

	for k := 2; k <= n; k++ {
		for i := 0; i+k <= n; i++ {
			j := i + k - 1
			dp[i][j] = k
			for a := i + 1; a <= j; a++ {
				tmp := dp[i][a-1] + dp[a][j]
				if s[i] == s[a] {
					tmp--
				}
				if tmp < dp[i][j] {
					dp[i][j] = tmp
				}
			}
		}
	}
	return dp[0][n-1]
}

func main() {
	fmt.Println(strangePrinter("aaabbb"))
	fmt.Println(strangePrinter("aba"))
	fmt.Println(strangePrinter("abac"))
	fmt.Println(strangePrinter("abaca"))

}
