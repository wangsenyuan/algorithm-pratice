package p2615

import "sort"

func distance(nums []int) []int64 {
	n := len(nums)

	type Pair struct {
		first  int
		second int
	}

	P := make([]Pair, n)
	for i := 0; i < n; i++ {
		P[i] = Pair{nums[i], i}
	}

	sort.Slice(P, func(i, j int) bool {
		return P[i].first < P[j].first || P[i].first == P[j].first && P[i].second < P[j].second
	})

	res := make([]int64, n)

	for i := 0; i < n; {
		j := i
		var sum int64
		for i < n && P[i].first == P[j].first {
			sum += int64(P[i].second)
			i++
		}
		var pref int64
		for k := j; k < i; k++ {
			res[P[k].second] = sum - int64(i-k)*int64(P[k].second)
			res[P[k].second] += int64(k-j)*int64(P[k].second) - pref

			sum -= int64(P[k].second)
			pref += int64(P[k].second)
		}

	}

	return res
}
