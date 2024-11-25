package p3361

const inf = 1 << 60

func shiftDistance(s string, t string, nextCost []int, previousCost []int) int64 {
	var dp [26][26]int
	for i := 0; i < 26; i++ {
		for j := 0; j < 26; j++ {
			dp[i][j] = inf
		}
		dp[i][i] = 0
	}

	for i := 0; i < 26; i++ {
		var sum int
		for j := 1; j < 26; j++ {
			sum += nextCost[(i+j-1)%26]
			dp[i][(i+j)%26] = min(dp[i][(i+j)%26], sum)
		}
		sum = 0
		for j := 1; j < 26; j++ {
			sum += previousCost[(i-(j-1)+26)%26]
			dp[i][(i-j+26)%26] = min(dp[i][(i-j+26)%26], sum)
		}
	}

	var ans int

	for i := range len(s) {
		x := int(s[i] - 'a')
		y := int(t[i] - 'a')
		ans += dp[x][y]
	}
	return int64(ans)
}
