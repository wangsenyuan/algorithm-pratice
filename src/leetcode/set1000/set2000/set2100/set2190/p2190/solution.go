package p2190

import "sort"

func sortJumbled(mapping []int, nums []int) []int {

	n := len(nums)

	P := make([]Pair, n)

	for i := 0; i < n; i++ {
		P[i] = Pair{transform(nums[i], mapping), i}
	}

	sort.Slice(P, func(i, j int) bool {
		return P[i].first < P[j].first || P[i].first == P[j].first && P[i].second < P[j].second
	})

	res := make([]int, n)

	for i := 0; i < n; i++ {
		res[i] = nums[P[i].second]
	}

	return res
}

func transform(num int, mapping []int) int64 {
	if num == 0 {
		return int64(mapping[0])
	}
	var base int64 = 1
	var res int64

	for num > 0 {
		cur := num % 10
		num /= 10
		cur = mapping[cur]
		res += int64(cur) * base
		base *= 10
	}
	return res
}

type Pair struct {
	first  int64
	second int
}
