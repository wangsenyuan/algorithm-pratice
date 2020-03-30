package p1394

import "sort"

func numTeams1(rating []int) int {
	n := len(rating)

	var res int

	for i := 1; i < n-1; i++ {
		var cnt1, cnt2 int
		for j := 0; j < i; j++ {
			if rating[j] < rating[i] {
				cnt1++
			} else if rating[j] > rating[i] {
				cnt2++
			}
		}

		var cnt3, cnt4 int

		for j := n - 1; j > i; j-- {
			if rating[j] > rating[i] {
				cnt3++
			} else if rating[j] < rating[i] {
				cnt4++
			}
		}

		res += cnt1*cnt3 + cnt2*cnt4
	}

	return res
}

func numTeams(rating []int) int {
	n := len(rating)
	ps := make([]*Pair, n)

	for i := 0; i < n; i++ {
		ps[i] = &Pair{rating[i], i}
	}

	sort.Sort(Pairs(ps))

	// n = j

	cnt1 := make([]int, n)
	cnt2 := make([]int, n)
	cnt3 := make([]int, n)
	cnt4 := make([]int, n)
	bit := CreateBIT(n)

	for i := 0; i < n; i++ {
		p := ps[i]
		cnt2[p.second] = bit.Get(p.second)
		cnt4[p.second] = bit.Get(n-1) - cnt2[p.second]

		bit.Update(p.second, 1)
	}

	bit = CreateBIT(n)

	for i := n - 1; i >= 0; i-- {
		p := ps[i]

		cnt1[p.second] = bit.Get(p.second)
		cnt3[p.second] = bit.Get(n-1) - cnt1[p.second]
		bit.Update(p.second, 1)
	}
	var res int
	for i := 0; i < n; i++ {
		res += cnt1[i]*cnt4[i] + cnt2[i]*cnt3[i]
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

type BIT struct {
	arr  []int
	size int
}

func CreateBIT(n int) BIT {
	arr := make([]int, n+1)
	return BIT{arr, n}
}

func (bit *BIT) Update(pos int, val int) {
	pos++
	for pos <= bit.size {
		bit.arr[pos] += val
		pos += pos & -pos
	}
}

func (bit *BIT) Get(pos int) int {
	var res int
	pos++
	for pos > 0 {
		res += bit.arr[pos]
		pos -= pos & -pos
	}
	return res
}
