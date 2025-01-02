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
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = -1
	}

	check := func(l int, r int) bool {
		if dp[r] >= l {
			return true
		}
		for j := r - 1; j > l; j-- {
			k := r - j
			if j-k < l {
				return false
			}
			if s[j-k] == s[j] && s[j] == s[r] {
				dp[r] = j - k
				return true
			}
		}
		return false
	}
	var ans int
	for l, r := 0, 1; l < n; l++ {
		for r < n && !check(l, r) {
			r++
		}
		ans += n - r
	}
	return ans
}
