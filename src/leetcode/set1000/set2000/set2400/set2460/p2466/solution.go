package p2466

const MOD = 1000000007

func modAdd(a, b int) int {
	a += b
	if a >= MOD {
		a -= MOD
	}
	return a
}

func countGoodStrings(low int, high int, zero int, one int) int {
	dp := make([]int, high+1)
	dp[0] = 1
	var res int
	for i := 1; i <= high; i++ {
		if i-zero >= 0 {
			dp[i] = modAdd(dp[i], dp[i-zero])
		}
		if i-one >= 0 {
			dp[i] = modAdd(dp[i], dp[i-one])
		}
		if i >= low {
			res = modAdd(res, dp[i])
		}
	}
	return res
}
