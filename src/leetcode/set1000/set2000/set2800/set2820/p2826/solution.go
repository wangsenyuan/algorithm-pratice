package p2826

func minimumOperations(nums []int) int {
	//n := len(nums)
	// dp[1] = 到i为止都是1
	dp := make([]int, 4)
	fp := make([]int, 4)

	for _, num := range nums {
		if num == 3 {
			fp[3] = min(dp[1:]...)
		} else {
			fp[3] = min(dp[1:]...) + 1
		}
		if num == 2 {
			fp[2] = min(dp[1], dp[2])
		} else {
			fp[2] = min(dp[1], dp[2]) + 1
		}
		if num == 1 {
			fp[1] = dp[1]
		} else {
			fp[1] = dp[1] + 1
		}
		copy(dp, fp)
	}
	return min(dp[1:]...)
}

func min(arr ...int) int {
	res := arr[0]
	for i := 1; i < len(arr); i++ {
		if res > arr[i] {
			res = arr[i]
		}
	}
	return res
}
