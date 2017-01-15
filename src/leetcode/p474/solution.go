package main

import "fmt"

func main() {
	strs := []string{"10", "0001", "111001", "1", "0"}
	m := 5
	n := 3
	fmt.Println(findMaxForm(strs, m, n))
}

func findMaxForm(strs []string, m int, n int) int {
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for _, str := range strs {
		a, b := countZeroOne(str)

		for i := m; i >= a; i-- {
			for j := n; j >= b; j-- {
				if 1+dp[i-a][j-b] > dp[i][j] {
					dp[i][j] = 1 + dp[i-a][j-b]
				}
			}
		}
	}

	return dp[m][n]
}

func countZeroOne(str string) (int, int) {
	zeros, ones := 0, 0
	for i := range str {
		if str[i] == '0' {
			zeros++
		} else {
			ones++
		}
	}

	return zeros, ones
}
