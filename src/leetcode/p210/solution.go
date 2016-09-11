package main

func minSubArrayLen(s int, nums []int) int {
	res := len(nums) + 1
	sum := 0

	for i, j := 0, 0; i < len(nums); i++ {
		for j < len(nums) && sum < s {
			sum += nums[j]
			j++
		}
		if sum < s {
			break
		}
		if j-i < res {
			res = j - i
		}
		sum -= nums[i]
	}

	if res == len(nums)+1 {
		return 0
	}
	return res
}
