package p1514

import "sort"

const MOD = 1000000007

func rangeSum(nums []int, n int, left int, right int) int {
	sums := make([]int64, 0, n*(n+1)/2)

	for i := 0; i < n; i++ {
		var sum int64
		for j := i; j < n; j++ {
			sum += int64(nums[j])
			sums = append(sums, sum)
		}
	}

	sort.Sort(Int64s(sums))

	var res int64

	for i := left - 1; i < right; i++ {
		res += sums[i]
		res %= MOD
	}

	return int(res)
}

type Int64s []int64

func (this Int64s) Len() int {
	return len(this)
}

func (this Int64s) Less(i, j int) bool {
	return this[i] < this[j]
}

func (this Int64s) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}
