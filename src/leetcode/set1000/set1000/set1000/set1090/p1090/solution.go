package p1090

import "sort"

func largestValsFromLabels(values []int, labels []int, num_wanted int, use_limit int) int {
	n := len(values)
	ps := make([]Pair, n)

	for i := 0; i < n; i++ {
		ps[i] = Pair{values[i], labels[i]}
	}

	sort.Sort(PS(ps))
	var res int
	used := make(map[int]int)

	for i := 0; i < n && num_wanted > 0; i++ {
		p := ps[i]
		v := p.first
		l := p.second
		if used[l] < use_limit {
			res += v
			used[l]++
			num_wanted--
		}
	}

	return res
}

type Pair struct {
	first  int
	second int
}

type PS []Pair

func (this PS) Len() int {
	return len(this)
}

func (this PS) Less(i, j int) bool {
	return this[i].first > this[j].first
}

func (this PS) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}
