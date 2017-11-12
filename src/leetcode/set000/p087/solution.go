package main

import "fmt"

func main() {
	fmt.Println(isScramble("rgtae", "great"))
	fmt.Println(isScramble("abc", "bca"))
}

func isScramble(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	var check func(i, j, l int) bool

	check = func(i, j, l int) bool {
		cs := make(map[byte]int)
		for k := 0; k < l; k++ {
			cs[s1[i+k]]++
			cs[s2[j+k]]--
		}

		for _, v := range cs {
			if v != 0 {
				return false
			}
		}
		if l == 1 {
			return true
		}
		for k := 1; k < l; k++ {
			if k <= l-k && (check(i, j, k) && check(i+k, j+k, l-k) || check(i, j+l-k, k) && check(i+k, j, l-k)) {
				return true
			} else if k > l-k && (check(i, j+l-k, k) && check(i+k, j, l-k) || check(i, j, k) && check(i+k, j+k, l-k)) {
				return true
			}
		}

		return false
	}

	return check(0, 0, len(s1))
}

func isScramble2(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	n := len(s1)
	dp := make([][][]bool, n)
	for i := range dp {
		dp[i] = make([][]bool, n)
		for j := range dp[i] {
			dp[i][j] = make([]bool, n+1)
			dp[i][j][0] = true
			dp[i][j][1] = s1[i] == s2[j]
		}
	}

	for i := n - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			for k := 2; k <= min(n-i, n-j); k++ {
				for l := 1; l < k; l++ {
					dp[i][j][k] = dp[i][j][k] || (dp[i][j][l] && dp[i+l][j+l][k-l])
					dp[i][j][k] = dp[i][j][k] || (dp[i][j+k-l][l] && dp[i+l][j][k-l])
					if dp[i][j][k] {
						break
					}
				}
			}
		}
	}

	return dp[0][0][n]
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
