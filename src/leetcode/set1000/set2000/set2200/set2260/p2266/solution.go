package p2266

const MOD = 1000000007

func countTexts(pressedKeys string) int {
	// 同一个数字最多出现4次
	n := len(pressedKeys)

	dp := make([]int, n+1)
	dp[0] = 1
	for i := 0; i < n; i++ {
		for j := 1; i+j <= n; j++ {
			if pressedKeys[i+j-1] != pressedKeys[i] {
				break
			}
			if pressedKeys[i] == '7' || pressedKeys[i] == '9' {
				// at most 4 times for 7 & 9
				if j > 4 {
					break
				}
			} else {
				if j > 3 {
					break
				}
			}
			dp[i+j] += dp[i]
			if dp[i+j] >= MOD {
				dp[i+j] -= MOD
			}
		}
	}
	return dp[n]
}
