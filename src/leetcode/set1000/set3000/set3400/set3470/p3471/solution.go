package p3471

import "slices"

func largestInteger(nums []int, k int) int {
	n := len(nums)
	x := slices.Max(nums)

	if k == n {
		return x
	}
	cnt := make([]int, x+1)
	for _, num := range nums {
		cnt[num]++
	}
	if k == 1 {
		// 最大的出现一次的数
		for y := x; y >= 0; y-- {
			if cnt[y] == 1 {
				return y
			}
		}
	} else {
		// k < n, 只需要检查两头的数
		var arr []int
		if cnt[nums[0]] == 1 {
			arr = append(arr, nums[0])
		}
		if cnt[nums[n-1]] == 1 {
			arr = append(arr, nums[n-1])
		}
		if len(arr) > 0 {
			return slices.Max(arr)
		}
	}

	return -1
}
