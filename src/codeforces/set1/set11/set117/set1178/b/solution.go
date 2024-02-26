package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	s := readString(reader)
	res := solve(s)
	fmt.Println(res)
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
}

func solve(s string) int {
	n := len(s)
	// dp[0] for first vv
	// dp[1] for vvo
	dp := make([]int, 2)

	var res int
	for i := 0; i < n; i++ {
		if s[i] == 'o' {
			dp[1] += dp[0]
		} else {
			// s[i] = 'v'
			if i+1 < n && s[i+1] == 'v' {
				res += dp[1]
				dp[0]++
			}
		}
	}

	return res
}
