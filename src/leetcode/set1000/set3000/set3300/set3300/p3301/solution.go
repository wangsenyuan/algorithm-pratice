package p3301

import "sort"

func maximumTotalSum(maximumHeight []int) int64 {
	sort.Ints(maximumHeight)
	n := len(maximumHeight)
	res := maximumHeight[n-1]
	next := maximumHeight[n-1]
	for i := n - 2; i >= 0; i-- {
		cur := min(maximumHeight[i], next-1)
		if cur == 0 {
			return -1
		}
		res += cur
		next = cur
	}
	return int64(res)
}
