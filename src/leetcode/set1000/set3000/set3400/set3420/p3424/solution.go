package p3424

import "sort"

func minCost(arr []int, brr []int, k int64) int64 {
	var res1 int64
	n := len(arr)
	for i := range n {
		res1 += int64(abs(arr[i] - brr[i]))
	}

	sort.Ints(arr)
	sort.Ints(brr)
	res2 := k
	for i := range n {
		res2 += int64(abs(arr[i] - brr[i]))
	}
	return min(res1, res2)
}

func abs(num int) int {
	return max(num, -num)
}
