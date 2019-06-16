package p300

import "sort"

func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	n := len(nums)

	ord := make([]int, n)
	sz := 0
	for i := 0; i < n; i++ {
		x := nums[i]
		j := sort.Search(sz, func(j int) bool {
			return ord[j] >= x
		})
		if j == sz {
			ord[sz] = x
			sz++
		} else {
			ord[j] = x
		}
	}

	return sz
}
