package lcp14

const A = 1000010

var spf [A]int

func init() {
	for x := 2; x < A; x++ {
		if spf[x] == 0 {
			for y := x; y < A; y += x {
				if spf[y] == 0 {
					spf[y] = x
				}
			}
		}
	}
}

func splitArray(nums []int) int {
	pos := make(map[int]int)
	n := len(nums)
	dp := make([]int, n+1)

	for i := 1; i <= n; i++ {
		var best = dp[i-1] + 1

		cur := nums[i-1]

		for cur > 1 {
			if j, found := pos[spf[cur]]; found {
				best = min(best, dp[j-1]+1)
			}
			cur = cur / spf[cur]
		}

		dp[i] = best

		cur = nums[i-1]
		for cur > 1 {
			if j, found := pos[spf[cur]]; !found || dp[j-1] > dp[i-1] {
				pos[spf[cur]] = i
			}
			cur /= spf[cur]
		}
	}

	return dp[n]
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
