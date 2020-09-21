package p1591

import "sort"

const MOD = 1000000007

func maxSumRangeQuery(nums []int, requests [][]int) int {
	n := len(nums)
	cnt := make([]int, n+1)
	for _, req := range requests {
		u, v := req[0], req[1]
		cnt[u]++
		cnt[v+1]--
	}

	var sum int
	ps := make([]Pair, n)
	for i := 0; i < n; i++ {
		sum += cnt[i]
		ps[i] = Pair{sum, i}
	}

	sort.Sort(Pairs(ps))
	sort.Ints(nums)
	arr := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		arr[ps[i].second] = nums[i]
	}
	sum = 0
	pref := make([]int, n+1)
	for i := 0; i < n; i++ {
		sum += arr[i]
		if sum >= MOD {
			sum -= MOD
		}
		pref[i+1] = sum
	}

	var res int

	for _, req := range requests {
		u, v := req[0], req[1]
		res += pref[v+1] - pref[u]
		if res < 0 {
			res += MOD
		}

		if res >= MOD {
			res -= MOD
		}
	}

	return res
}

type Pair struct {
	first, second int
}

type Pairs []Pair

func (ps Pairs) Len() int {
	return len(ps)
}

func (ps Pairs) Less(i, j int) bool {
	return ps[i].first < ps[j].first || ps[i].first == ps[j].first && ps[i].second < ps[j].second
}

func (ps Pairs) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}
