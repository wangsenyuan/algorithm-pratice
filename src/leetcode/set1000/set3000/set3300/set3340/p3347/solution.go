package p3347

import (
	"sort"
)

func maxFrequency(nums []int, k int, numOperations int) int {
	sort.Ints(nums)
	// 因为可以选择0， 所以，相当于可以不操作某些数
	// 就是找出重叠区间最大的那些
	// 但是有个问题是，操作的次数有上限
	n := len(nums)

	// 两种情况，一种是，中间的数，不变，两边的数，变化到它这里
	// 从中选择 numop个数
	// 还有一个就是，(a, b)， a前面的变大，b后面的变小
	// 看看有多少个数可以重叠
	var best int

	for l, i, r := 0, 0, 0; i < n; {
		j := i
		for l < j && nums[j]-nums[l] > k {
			l++
		}
		r = max(r, i)
		for r < n && nums[r]-nums[j] <= k {
			r++
		}

		for i < n && nums[i] == nums[j] {
			i++
		}

		cnt := i - j
		side := j - l + r - i
		best = max(best, cnt+min(side, numOperations))
	}

	// 然后就是所有的数，都变化的情况
	for l, r := 0, 0; r < n; r++ {
		for l < r && nums[r]-nums[l] > 2*k {
			l++
		}
		best = max(best, min(numOperations, r-l+1))
	}

	return best
}
