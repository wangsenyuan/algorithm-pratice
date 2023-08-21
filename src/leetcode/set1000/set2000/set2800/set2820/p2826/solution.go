package p2826

func minimumOperations(nums []int) int {
	//n := len(nums)
	// dp[1] = 到i为止都是1
	dp := make([]int, 4)
	//fp := make([]int, 4)

	for _, num := range nums {
		dp[3] = min(dp[1:]...)
		if num != 3 {
			dp[3]++
		}
		dp[2] = min(dp[1], dp[2])
		if num != 2 {
			dp[2]++
		}
		if num != 1 {
			dp[1]++
		}
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
