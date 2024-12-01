package p3372

import "sort"

func getLargestOutlier(nums []int) int {
	sort.Ints(nums)
	// x + x + y = sum

	n := len(nums)
	var sum int

	for _, num := range nums {
		sum += num
	}

	for i := n - 1; i >= 0; i-- {
		y := nums[i]
		x := sum - y
		if x%2 == 0 {
			x /= 2
			// 如果同时存在前缀和x，和数字x
			j := sort.SearchInts(nums, x)
			if j == n || nums[j] > x {
				j--
			}
			if j == i {
				j--
			}
			if j < 0 || nums[j] != x {
				continue
			}
			// nums[j] = x, 剩下的肯定和肯定是x
			return y
		}
	}

	return 0
}
