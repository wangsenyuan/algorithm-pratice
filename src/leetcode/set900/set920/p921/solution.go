package p921

func minAddToMakeValid(S string) int {
	n := len(S)
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = -1
	}
	// dp[i] = j, means the far pos that make [i..j] balance

	for i := 0; i < n; i++ {
		if S[i] == ')' {
			continue
		}
		var level int
		for j := i; j < n; j++ {
			if S[j] == '(' {
				level++
			} else if S[j] == ')' {
				level--
			}
			if level == 0 {
				dp[i] = j
				break
			}
		}
	}

	var ans int
	var pos int
	for pos < n {
		if dp[pos] < 0 {
			ans++
			pos++
		} else {
			pos = dp[pos] + 1
		}
	}
	return ans
}
