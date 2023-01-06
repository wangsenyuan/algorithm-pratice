package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	s, _ := reader.ReadString('\n')
	a, b := solve(s)

	fmt.Printf("%d %d\n", a, b)
}

func solve(s string) (int, int) {
	n := len(s)
	dp := make([]int, n)

	st := make([]int, n)
	var p int

	for i := 0; i < n; i++ {
		dp[i] = -1
		if s[i] == '(' {
			st[p] = i
			p++
			continue
		}
		if s[i] == ')' {
			if p == 0 {
				continue
			}
			dp[i] = st[p-1]
			p--
			if dp[i] > 0 && dp[dp[i]-1] >= 0 {
				dp[i] = dp[dp[i]-1]
			}
		}
	}

	var best int
	var cnt = 1
	for i := 0; i < n; i++ {
		if dp[i] < 0 {
			continue
		}
		if i-dp[i]+1 > best {
			best = i - dp[i] + 1
			cnt = 1
		} else if i-dp[i]+1 == best {
			cnt++
		}
	}

	return best, cnt
}
