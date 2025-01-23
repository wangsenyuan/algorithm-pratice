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

const inf = 1 << 60

func solve(s string) int {
	n := len(s)
	dp := make([]int, n+1)
	dp[0] = 0
	for i := 0; i < n; i++ {
		dp[i+1] = dp[i]
		var sum int
		for j := 0; j < 3 && i-j >= 0; j++ {
			x := int(s[i-j] - '0')
			sum += x
			sum %= 3
			if j > 0 && s[i-j] == '0' {
				// leading zero
				continue
			}
			if sum == 0 {
				dp[i+1] = max(dp[i+1], dp[i-j]+1)
			}
		}
	}
	return dp[n]
}
