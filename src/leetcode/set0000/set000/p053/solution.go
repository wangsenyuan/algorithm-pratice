package p053

import "math"

func maxSubArray(nums []int) int {
	maxSum, currSum := math.MinInt32, 0

	for _, num := range nums {
		currSum += num
		if currSum > maxSum {
			maxSum = currSum
		}
		if currSum < 0 {
			currSum = 0
		}
	}

	return maxSum
}
