package main

import (
	"bufio"
	"fmt"
	"os"
)

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	s := readString(reader)

	res := solve(s)
	fmt.Println(res)
}

const T = 20

func solve(s string) int {

	TT := 1 << T
	// dp[state] = state的最左边的位置
	dp := make([]int, TT)

	n := len(s)

	var best int

	for i := 0; i < n; i++ {
		var state int
		for j := i; j >= 0; j-- {
			x := int(s[j] - 'a')
			if (state>>x)&1 == 1 {
				// duplicates not allowed
				break
			}
			state |= 1 << x
			dp[state] = max(dp[state], i-j+1)
			best = max(best, i-j+1)
		}
	}

	for state := 1; state < TT; state++ {
		for i := 0; i < T; i++ {
			if (state>>i)&1 == 1 {
				dp[state] = max(dp[state], dp[state^(1<<i)])
			}
		}
	}

	mask := TT - 1
	// do agein
	for i := 0; i < n; i++ {
		var state int
		for j := i; j >= 0; j-- {
			x := int(s[j] - 'a')
			if (state>>x)&1 == 1 {
				// duplicates not allowed
				break
			}
			state |= 1 << x
			best = max(best, i-j+1+dp[mask^state])
		}
	}
	return best
}
