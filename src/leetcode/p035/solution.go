package main

func searchInsert(nums []int, target int) int {
	i, j := 0, len(nums)
	for i < j {
		k := i + (j-i)/2
		if nums[k] == target {
			return k
		} else if nums[k] < target {
			i = k + 1
		} else {
			j = k
		}
	}

	return j
}
