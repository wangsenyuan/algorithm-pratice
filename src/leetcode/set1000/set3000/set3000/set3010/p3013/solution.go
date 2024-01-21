package p3013

import (
	"math/bits"
	"sort"
)

func canSortArray(nums []int) bool {
	n := len(nums)

	for i := 0; i < n; {
		j := i
		cnt := bits.OnesCount(uint(nums[j]))
		for i < n && bits.OnesCount(uint(nums[i])) == cnt {
			i++
		}
		sort.Ints(nums[j:i])
	}

	return sort.IntsAreSorted(nums)
}
