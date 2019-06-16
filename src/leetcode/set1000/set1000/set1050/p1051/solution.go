package p1051

import "sort"

func heightChecker(heights []int) int {
	n := len(heights)
	ps := make(Pairs, n)
	for i := 0; i < n; i++ {
		ps[i] = Pair{first: heights[i], second: i}
	}

	sort.Sort(ps)

	var same int
	for i := 0; i < n; i++ {
		h := heights[i]
		p := ps[i]
		if h == p.first {
			same++
		}
	}
	return n - same
}

type Pair struct {
	first  int
	second int
}

type Pairs []Pair

func (this Pairs) Len() int {
	return len(this)
}

func (this Pairs) Less(i, j int) bool {
	return this[i].first < this[j].first
}

func (this Pairs) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}
