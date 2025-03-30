package p3500

func minimumCost(nums []int, cost []int, k int) int64 {
	n := len(nums)
	pref := make([]int, n+1)
	for i, x := range cost {
		pref[i+1] = pref[i] + x
	}
	dp := make([]int, n+1)
	var sum int
	for i, x := range nums {
		sum += x
		i++
		var tmp int = 1e18
		for j := range i {
			tmp = min(tmp, dp[j]+k*(pref[n]-pref[j])+sum*(pref[i]-pref[j]))
		}
		dp[i] = tmp
	}

	return int64(dp[n])
}
