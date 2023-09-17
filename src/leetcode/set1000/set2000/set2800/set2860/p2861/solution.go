package p2861

import "sort"

func countWays(nums []int) int {
	sort.Ints(nums)
	// 如果 i 被选中，那么选中group的size > num[i]
	// 且 如果i未被选中,未选中的group的size < num[i]
	var res int
	n := len(nums)

	// 全部被选中时
	if nums[n-1] < n {
		res++
	}

	// 全部未被选中时
	if nums[0] > 0 {
		res++
	}

	for i := 0; i+1 < n; i++ {
		// nums[i] = 最后一个被选中的学生
		if nums[i] < i+1 && i+1 < nums[i+1] {
			res++
		}
	}

	return res
}
