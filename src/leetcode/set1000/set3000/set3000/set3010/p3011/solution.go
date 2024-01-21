package p3011

import "sort"

func minimumArrayLength(nums []int) int {
	var res int
	sort.Ints(nums)

	n := len(nums)
	var i int
	for i < n && nums[i] == 0 {
		res++
		i++
	}

	if i == n {
		return res
	}

	j := i
	for i < n && nums[i] == nums[j] {
		i++
	}

	cnt := i - j

	if i == n || cnt == 1 {
		return res + (cnt+1)/2
	}

	for i < n && nums[i]%nums[j] == 0 {
		i++
	}
	if i < n {
		return res + 1
	}
	// cnt > 1 and i < n
	return res + (cnt+1)/2
}
