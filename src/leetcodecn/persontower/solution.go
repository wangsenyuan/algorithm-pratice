package persontower

import "sort"

func bestSeqAtIndex(height []int, weight []int) int {
	n := len(height)
	people := make([]Pair, n)

	for i := 0; i < n; i++ {
		people[i] = Pair{height[i], weight[i]}
	}

	sort.Sort(Pairs(people))

	// [w1, c1], [w2, c2], [w3, c3]
	// w1 > w2 > w3
	// c1 < c2 < c3

	stack := make([]int, n)
	var p int

	for i := 0; i < n; i++ {
		cur := people[i]

		j := sort.Search(p, func(j int) bool {
			return stack[j] <= cur.second
		})

		if j == p {
			stack[p] = cur.second
			p++
		} else {
			stack[j] = cur.second
		}
	}

	return p
}

type Pair struct {
	first, second int
}

type Pairs []Pair

func (ps Pairs) Len() int {
	return len(ps)
}

func (ps Pairs) Less(i, j int) bool {
	return ps[i].first > ps[j].first || ps[i].first == ps[j].first && ps[i].second < ps[j].second
}

func (ps Pairs) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}
