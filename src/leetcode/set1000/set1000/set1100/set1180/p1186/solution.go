package p1186

const INF = 1 << 20

func maximumSum(arr []int) int {
	n := len(arr)
	dp := make([]int, n)
	fp := make([]int, n)
	var best int
	for i := 0; i < n; i++ {
		best += arr[i]
		dp[i] = arr[i]
		if i > 0 && arr[i]+dp[i-1] > dp[i] {
			dp[i] = arr[i] + dp[i-1]
		}
	}

	for i := n - 1; i >= 0; i-- {
		fp[i] = arr[i]
		if i < n-1 && arr[i]+fp[i+1] > fp[i] {
			fp[i] = arr[i] + fp[i+1]
		}
	}

	for i := 0; i < n; i++ {
		if arr[i] < 0 {
			a := -INF
			if i > 0 {
				a = dp[i-1]
			}
			b := -INF
			if i < n-1 {
				b = fp[i+1]
			}

			if a < 0 && b < 0 {
				best = max(best, arr[i])
			} else if (a < 0) != (b < 0) {
				best = max(best, max(a, b))
			} else {
				best = max(best, a+b)
			}
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
