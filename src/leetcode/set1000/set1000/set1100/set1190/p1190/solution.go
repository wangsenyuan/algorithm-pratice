package p1190

import (
	"bytes"
)

func reverseParentheses(s string) string {
	s = "#" + s + "#"
	n := len(s)
	fp := make([]int, n)
	dp := make([]int, n)
	sp := make([]int, n)
	var p int

	for i := 0; i < n; i++ {
		if s[i] == '(' {
			sp[p] = i
			p++
		} else if s[i] == ')' {
			fp[sp[p-1]] = i
			dp[i] = sp[p-1]
			p--
		}
	}

	var loop func(start, end, level int)

	var buf bytes.Buffer

	loop = func(start, end, level int) {
		if level%2 == 0 {
			// no need to reverse
			for i := start; i <= end; i++ {
				if s[i] == '(' {
					loop(i+1, fp[i]-1, level+1)
					i = fp[i]
				} else {
					buf.WriteByte(s[i])
				}
			}
		} else {
			for i := end; i >= start; i-- {
				if s[i] == ')' {
					loop(dp[i]+1, i-1, level+1)
					i = dp[i]
				} else {
					buf.WriteByte(s[i])
				}
			}
		}
	}

	loop(0, n-1, 0)
	res := buf.String()

	return res[1 : len(res)-1]
}
