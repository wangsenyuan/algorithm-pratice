package p1915

func wonderfulSubstrings(word string) int64 {
	M := 1 << 10
	dp := make([]int64, M)
	dp[0] = 1

	var res int64
	var cur int
	for i := 0; i < len(word); i++ {
		x := int(word[i] - 'a')
		cur ^= 1 << x

		res += dp[cur]

		for j := 0; j < 10; j++ {
			tmp := cur ^ (1 << j)
			res += dp[tmp]
		}

		dp[cur]++
	}

	return res
}
