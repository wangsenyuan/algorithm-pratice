package p3047

import "sort"

func earliestSecondToMarkIndices(nums []int, changeIndices []int) int {
	n := len(nums)
	m := len(changeIndices)
	last := make([]int, n)

	check := func(x int) bool {
		for i := 0; i < n; i++ {
			last[i] = -1
		}
		for i := 0; i < x; i++ {
			j := changeIndices[i]
			last[j-1] = i
		}
		for i := 0; i < n; i++ {
			if last[i] < 0 {
				return false
			}
		}

		var cnt int
		for i := 0; i < x; i++ {
			j := changeIndices[i] - 1

			if last[j] != i {
				cnt++
			} else {
				if cnt < nums[j] {
					return false
				}
				cnt -= nums[j]
			}
		}

		return cnt >= 0
	}

	res := sort.Search(m+1, check)
	if res > m {
		return -1
	}
	return res
}
