package p1054

import (
	"sort"
)

func rearrangeBarcodes(barcodes []int) []int {
	n := len(barcodes)

	cnt := make(map[int]int)

	for i := 0; i < n; i++ {
		cnt[barcodes[i]]++
	}
	ps := make(Pairs, len(cnt))
	var idx int
	for k, v := range cnt {
		ps[idx] = &Pair{v, k}
		idx++
	}

	sort.Sort(sort.Reverse(ps))

	res := make([]int, n)

	var j int
	var i int
	for j < n {
		p := ps[i]
		c := p.first
		num := p.second
		for c > 0 && j < n {
			res[j] = num
			j += 2
			c--
		}
		if c == 0 {
			i++
		} else {
			p.first = c
		}
	}
	j = 1
	for j < n {
		p := ps[i]
		c := p.first
		num := p.second
		for c > 0 && j < n {
			res[j] = num
			j += 2
			c--
		}
		if c == 0 {
			i++
		}
	}

	return res
}

type Pair struct {
	first  int
	second int
}

type Pairs []*Pair

func (this Pairs) Len() int {
	return len(this)
}

func (this Pairs) Less(i, j int) bool {
	return this[i].first < this[j].first
}

func (this Pairs) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}
