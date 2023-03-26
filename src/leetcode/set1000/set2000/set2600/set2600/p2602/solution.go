package p2602

import "sort"

func minOperations(nums []int, queries []int) []int64 {
	sort.Ints(nums)
	n := len(queries)
	qs := make([]Pair, n)

	for i := 0; i < n; i++ {
		qs[i] = Pair{queries[i], i}
	}

	sort.Slice(qs, func(i, j int) bool {
		return qs[i].first < qs[j].first
	})

	var right int64
	for _, num := range nums {
		right += int64(num)
	}
	var left int64

	m := len(nums)

	ans := make([]int64, n)
	for i, j := 0, 0; i < n; i++ {
		x := qs[i].first
		for j < m && nums[j] <= x {
			left += int64(nums[j])
			right -= int64(nums[j])
			j++
		}
		a := int64(x)*int64(j) - left
		b := right - int64(x)*int64(m-j)
		ans[qs[i].second] = a + b
	}
	return ans
}

type Pair struct {
	first  int
	second int
}
