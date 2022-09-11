package p2404

import "sort"

func mostFrequentEven(nums []int) int {
	sort.Ints(nums)

	n := len(nums)

	ans := -1
	best := 0

	for i := 0; i < n; {
		j := i
		for i < n && nums[i] == nums[j] {
			i++
		}
		cnt := i - j
		if nums[j]%2 == 0 {
			if ans < 0 || best < cnt {
				ans = nums[j]
				best = cnt
			}
		}
	}

	return ans
}
