package main

func removeDuplicates(nums []int) int {
	p := 1

	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[p-1] && (p >= 2 && nums[p-1] == nums[p-2]) {
			continue
		}
		nums[p] = nums[i]
		p++
	}

	return p
}
