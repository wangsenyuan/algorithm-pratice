package p3448

func countSubstrings(s string) int64 {
	var dp [10][10]int

	var res int
	for _, x := range []byte(s) {
		v := int(x - '0')
		for j := 1; j < 10; j++ {
			var ndp [10]int
			ndp[v%j]++
			for r, w := range dp[j] {
				ndp[(r*10+v)%j] += w
			}
			dp[j] = ndp
		}
		res += dp[v][0]
	}

	return int64(res)
}
