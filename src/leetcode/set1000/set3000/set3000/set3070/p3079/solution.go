package p3079

import "sort"

func maximumHappinessSum(happiness []int, k int) int64 {
	sort.Ints(happiness)

	// 越大的越早选
	var res int
	n := len(happiness)
	var dec int
	for i := 0; i < k; i++ {
		cur := max(0, happiness[n-1-i]-dec)
		res += cur
		dec++
	}
	return int64(res)
}
