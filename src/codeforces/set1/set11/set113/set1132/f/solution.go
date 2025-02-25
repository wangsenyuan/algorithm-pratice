package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	readString(reader)
	s := readString(reader)
	res := solve(s)
	fmt.Println(res)
}

func readString(r *bufio.Reader) string {
	s, _ := r.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
}

func solve(s string) int {
	n := len(s)

	dp := make([][]int, n)
	for i := range n {
		dp[i] = make([]int, n)
		for j := range n {
			dp[i][j] = -1
		}
	}

	var dfs func(l int, r int) int
	dfs = func(l int, r int) int {
		if l > r {
			return 0
		}
		if l == r {
			return 1
		}
		if dp[l][r] >= 0 {
			return dp[l][r]
		}
		dp[l][r] = 1 + dfs(l+1, r)

		for i := l + 1; i <= r; i++ {
			if s[i] == s[l] {
				dp[l][r] = min(dp[l][r], dfs(l+1, i-1)+dfs(i, r))
			}
		}

		return dp[l][r]
	}

	return dfs(0, n-1)
}
