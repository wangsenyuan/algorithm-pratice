package p2272

const INF = 1 << 28

func largestVariance(s string) int {
	n := len(s)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, 26*26)
	}

	cnt := make([]int, 26)
	prev := make([]int, 26)
	for i := 0; i < 26; i++ {
		prev[i] = -1
	}

	var best int

	for i := 0; i < n; i++ {
		x := int(s[i] - 'a')
		cnt[x]++

		for j := 0; j < 26; j++ {
			if j == x || prev[j] < 0 {
				continue
			}
			k := prev[j]
			tmp := cnt[x] - cnt[j]
			if k > 0 {
				tmp -= dp[k-1][x*26+j]
			}
			best = max(best, tmp)

			tmp = cnt[j] - cnt[x]
			if k > 0 {
				tmp -= dp[k-1][j*26+x]
			}
			best = max(best, tmp)
		}
		if i > 0 {
			copy(dp[i], dp[i-1])
		}
		for j := 0; j < 26; j++ {
			if j == x {
				continue
			}
			dp[i][x*26+j] = min(cnt[x]-cnt[j], dp[i][x*26+j])
			dp[i][j*26+x] = min(cnt[j]-cnt[x], dp[i][j*26+x])
		}

		prev[x] = i
	}
	return best
}

func largestVariance1(s string) int {
	n := len(s)
	cnt := make([][]int, n)
	for i := 0; i < n; i++ {
		cnt[i] = make([]int, 26)
	}
	var best int
	for i := 0; i < n; i++ {
		if i > 0 {
			copy(cnt[i], cnt[i-1])
		}
		x := int(s[i] - 'a')
		cnt[i][x]++
		for j := 0; j < i; j++ {
			var tmp int
			for y := 0; y < 26; y++ {
				if x != y && cnt[i][y] > 0 {
					if j > 0 {
						tmp = max(tmp, cnt[i][x]-cnt[j-1][x]-(cnt[i][y]-cnt[j-1][y]))
						tmp = max(tmp, (cnt[i][y]-cnt[j-1][y])-(cnt[i][x]-cnt[j-1][x]))
					} else {
						tmp = max(tmp, cnt[i][x]-cnt[i][y])
						tmp = max(tmp, cnt[i][y]-cnt[i][x])
					}
				}
			}

			best = max(best, tmp)
		}
	}
	return best
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
