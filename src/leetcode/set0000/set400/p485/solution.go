package main

func findMaxConsecutiveOnes(nums []int) int {
	var res int
	var sum int
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if nums[i] == 0 {
			res = max(res, sum)
			sum = 0
		}
	}
	res = max(res, sum)
	return res
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
