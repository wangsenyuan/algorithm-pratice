package main

import "fmt"

func main() {
	fmt.Println(isMatch("aa", "a"))
	fmt.Println(isMatch("aa", "aa"))
	fmt.Println(isMatch("aaa", "aa"))
	fmt.Println(isMatch("aa", "a*"))
	fmt.Println(isMatch("aa", ".*"))
	fmt.Println(isMatch("ab", ".*"))
	fmt.Println(isMatch("aab", "c*a*b"))
}

func isMatch(s string, p string) bool {
	dp := make([][]bool, len(s)+1, len(s)+1)
	for i := range dp {
		dp[i] = make([]bool, len(p)+1, len(p)+1)
	}

	dp[0][0] = true

	for i := 1; i < len(p); i++ {
		if p[i] == '*' && dp[0][i-1] {
			dp[0][i+1] = true
		}
	}

	for i := 0; i < len(s); i++ {
		for j := 0; j < len(p); j++ {
			if p[j] == '.' || s[i] == p[j] {
				dp[i+1][j+1] = dp[i][j]
				continue
			}

			if p[j] == '*' {
				if s[i] != p[j-1] && p[j-1] != '.' {
					dp[i+1][j+1] = dp[i+1][j-1]
				} else if s[i] == p[j-1] || p[j-1] == '.' {
					dp[i+1][j+1] = dp[i][j+1] || dp[i+1][j] || dp[i+1][j-1]
				}
			}
		}
	}

	return dp[len(s)][len(p)]
}
