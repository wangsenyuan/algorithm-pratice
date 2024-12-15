package p1338

import "sort"

func minSetSize(arr []int) int {
	n := len(arr)

	freq := make(map[int]int)

	for _, x := range arr {
		freq[x]++
	}

	var nums []int
	for _, v := range freq {
		nums = append(nums, v)
	}

	sort.Ints(nums)

	var sum int

	for i := len(nums) - 1; i >= 0; i-- {
		sum += nums[i]
		if sum*2 >= n {
			return len(nums) - i
		}
	}
	return len(nums)
}
