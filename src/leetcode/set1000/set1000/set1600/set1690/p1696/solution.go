package p1696

func maxResult(nums []int, k int) int {
	n := len(nums)
	dp := make([]int, n)

	dp[0] = nums[0]

	que := make([]int, n)
	var front, end int
	que[end] = 0
	end++

	for i := 1; i < n; i++ {
		for front < end && que[front]+k < i {
			front++
		}
		dp[i] = dp[que[front]] + nums[i]
		for front < end && dp[que[end-1]] < dp[i] {
			end--
		}
		que[end] = i
		end++
	}

	return dp[n-1]
}
