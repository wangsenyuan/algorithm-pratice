package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	s, _ := reader.ReadString('\n')
	t, _ := reader.ReadString('\n')

	fmt.Println(solve(s[:len(s)-1], t[:len(t)-1]))
}

func solve(s, t string) int {
	n := len(s)
	m := len(t)
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, m+1)
		for j := 0; j <= m; j++ {
			dp[i][j] = n + 1
		}
	}

	nxt := make([]int, n)

	for i := 0; i < n; i++ {
		if s[i] != '.' {
			nxt[i] = -1
			var bal int
			for j := i; j < n; j++ {
				if s[j] == '.' {
					bal--
				} else {
					bal++
				}
				if bal == 0 {
					nxt[i] = j
					break
				}
			}
		}
	}

	dp[0][0] = 0

	for i := 0; i < n; i++ {
		for j := 0; j <= m; j++ {
			dp[i+1][j] = min(dp[i+1][j], dp[i][j]+1)
			if j < m && s[i] == t[j] {
				dp[i+1][j+1] = min(dp[i+1][j+1], dp[i][j])
			}
			if s[i] != '.' && nxt[i] != -1 {
				dp[nxt[i]+1][j] = min(dp[nxt[i]+1][j], dp[i][j])
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
