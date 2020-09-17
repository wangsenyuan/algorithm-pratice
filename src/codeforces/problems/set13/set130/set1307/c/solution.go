package main

import (
	"fmt"
)

func main() {
	// scanner := bufio.NewScanner(os.Stdin)

	// scanner.Scan()
	// s := scanner.Text()
	var s string
	fmt.Scanf("%s", &s)
	fmt.Println(solve(s))
}

func solve(s string) int64 {
	dp := make([]int64, 26)
	fp := make([][]int64, 26)

	for i := 0; i < 26; i++ {
		fp[i] = make([]int64, 26)
	}

	n := len(s)

	for i := n - 1; i >= 0; i-- {
		x := int(s[i] - 'a')

		for j := 0; j < 26; j++ {
			fp[x][j] += dp[j]
		}

		dp[x]++
	}

	var res int64
	for i := 0; i < 26; i++ {
		res = max(res, dp[i])
		for j := 0; j < 26; j++ {
			res = max(res, fp[i][j])
		}
	}

	return res
}

func max(a, b int64) int64 {
	if a >= b {
		return a
	}
	return b
}
