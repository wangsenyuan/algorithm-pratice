package contest

import "bytes"

func removeDuplicates(s string, k int) string {
	n := len(s)
	dp := make([]int, n)
	stack := make([]int, n)
	var p int
	for i := 0; i < n; i++ {
		if p > 0 && s[stack[p-1]] == s[i] {
			// can connect
			cnt := 1 + dp[p-1]
			r := cnt % k
			if r == 0 {
				// can concate and dimiss stack[p-1] & i
				p--
			} else {
				dp[p-1]++
				stack[p-1] = i
			}
		} else {
			stack[p] = i
			dp[p] = 1
			p++
		}
	}
	var buf bytes.Buffer

	for j := 0; j < p; j++ {
		i := stack[j]
		c := dp[j]
		for c > 0 {
			buf.WriteByte(s[i])
			c--
		}
	}
	return buf.String()
}
