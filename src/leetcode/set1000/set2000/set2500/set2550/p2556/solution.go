package p2556

import "sort"

func countFairPairs(nums []int, lower int, upper int) int64 {
	sort.Ints(nums)
	n := len(nums)
	var res int64
	for j := 0; j < n; j++ {
		// nums[i] + nums[j] >= lower i != j
		i := search(0, j, func(i int) bool {
			return nums[i]+nums[j] >= lower
		})
		if i == j {
			continue
		}
		// nums[i] + nums[j] >= lower
		// nums[i] + nums[j] <= upper
		k := search(i, j, func(k int) bool {
			return nums[k]+nums[j] > upper
		})
		if k <= i {
			continue
		}
		res += int64(k - i)
	}

	return res
}

func search(l int, r int, f func(int) bool) int {
	for l < r {
		mid := (l + r) / 2
		if f(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return r
}
