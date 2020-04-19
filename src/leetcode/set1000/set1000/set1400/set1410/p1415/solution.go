package p1415

import "bytes"

func getHappyString(n int, k int) string {
	dp := make([][]int, 3)
	for i := 0; i < 3; i++ {
		dp[i] = make([]int, n+1)
		dp[i][1] = 1
	}

	for j := 2; j <= n; j++ {
		for i := 0; i < 3; i++ {
			for ii := 0; ii < 3; ii++ {
				if i == ii {
					continue
				}
				dp[i][j] += dp[ii][j-1]
			}
		}
	}
	var pos int
	var buf bytes.Buffer
	var prev = -1
	for j := n; j > 0; j-- {
		var i int

		for i < 3 && (i == prev || pos+dp[i][j] < k) {
			if i == prev {
				i++
				continue
			}
			pos += dp[i][j]
			i++
		}

		if i == 3 || pos > k {
			return ""
		}
		buf.WriteByte(byte(i + 'a'))
		prev = i
	}

	return buf.String()
}
