package p1094

import "sort"

func carPooling(trips [][]int, capacity int) bool {
	n := len(trips)
	ps := make(PS, 2*n)
	for i := 0; i < n; i++ {
		trip := trips[i]
		a, b, c := trip[0], trip[1], trip[2]
		ps[2*i] = Pair{b, a}
		ps[2*i+1] = Pair{c, a * -1}
	}

	sort.Sort(ps)

	var took int

	for i := 0; i < len(ps); i++ {
		took += ps[i].second
		if took > capacity {
			return false
		}
	}

	return true
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
	return this[i].first < this[j].first || (this[i].first == this[j].first && this[i].second < this[j].second)
}

func (this PS) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}
