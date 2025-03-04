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
	ok := make([][]bool, n)
	for i := 0; i < n; i++ {
		ok[i] = make([]bool, n)
		ok[i][i] = true
	}

	for r := range n {
		for l := r - 1; l >= 0; l-- {
			if s[l] == s[r] {
				if l+1 == r || ok[l+1][r-1] {
					ok[l][r] = true
				}
			}
		}
	}

	// dp[i] 是到i为止有多少个回文
	dp := make([]int, n)

	for i := 0; i < n; i++ {
		for j := i; j >= 0; j-- {
			if ok[j][i] {
				dp[i]++
			}
		}
	}

	var res int
	var suf int
	for i := n - 1; i > 0; i-- {
		for j := i; j < n; j++ {
			if ok[i][j] {
				suf++
			}
		}
		res += dp[i-1] * suf
	}

	return res
}
