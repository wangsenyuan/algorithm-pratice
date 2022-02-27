package p2188

const LIMIT = 1 << 33

func minimumFinishTime(tires [][]int, changeTime int, numLaps int) int {
	n := len(tires)
	// dp[i] = 跑i圈所用的最小时间
	// dp[i+1] = dp[j] + changeTime + min_time_over( tiers user i + 1 - j 秒)
	// f + f * r + f * r * r ... = f * (1 + r + r * r)
	// f * ((1 - r**x) / (1 - r))

	X := make([]int64, 31)

	for i := 1; i <= 30; i++ {
		X[i] = LIMIT
		for j := 0; j < n; j++ {
			f, r := tires[j][0], tires[j][1]
			F, R := int64(f), int64(r)
			tmp := F * (1 - pow(R, i)) / (1 - R)
			if tmp > 0 && X[i] > tmp {
				X[i] = tmp
			}
		}
	}

	dp := make([]int64, numLaps+1)

	for i := 1; i <= numLaps; i++ {
		if i <= 30 {
			dp[i] = X[i]
		} else {
			dp[i] = LIMIT
		}
		for j := i - 1; j > 0 && i-j <= 30; j-- {
			dp[i] = min(dp[i], dp[j]+X[i-j]+int64(changeTime))
		}
	}
	return int(dp[numLaps])
}

func min(a, b int64) int64 {
	if a <= b {
		return a
	}
	return b
}

func pow(a int64, b int) int64 {
	var r int64 = 1
	for b > 0 {
		if b&1 == 1 {
			r *= a
		}
		a *= a
		b >>= 1
		if r >= LIMIT {
			break
		}
	}
	return r
}
