package p2593

import "sort"

func findScore(nums []int) int64 {
	type Pair struct {
		first  int
		second int
	}
	n := len(nums)
	p := make([]Pair, n)
	for i := 0; i < n; i++ {
		p[i] = Pair{nums[i], i}
	}
	sort.Slice(p, func(i, j int) bool {
		return p[i].first < p[j].first || p[i].first == p[j].first && p[i].second < p[j].second
	})
	marked := make([]bool, n)

	var res int64

	for _, cur := range p {
		if marked[cur.second] {
			continue
		}
		res += int64(cur.first)
		if cur.second > 0 {
			marked[cur.second-1] = true
		}
		if cur.second+1 < n {
			marked[cur.second+1] = true
		}
	}
	return res
}
