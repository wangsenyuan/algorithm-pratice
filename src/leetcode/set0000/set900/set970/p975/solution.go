package p975

import "sort"

func oddEvenJumps(A []int) int {
	n := len(A)

	ps := make(PS, n)

	for i := 0; i < n; i++ {
		ps[i] = P{A[i], i}
	}

	sort.Sort(ps)

	st := make([]int, n)
	var p int
	gt := make([]int, n)

	for i := 0; i < n; i++ {
		for p > 0 && st[p-1] < ps[i].second {
			gt[st[p-1]] = ps[i].second
			p--
		}
		st[p] = ps[i].second
		p++
	}

	sort.Sort(RPS(ps))
	p = 0
	lt := make([]int, n)

	for i := 0; i < n; i++ {
		for p > 0 && st[p-1] < ps[i].second {
			lt[st[p-1]] = ps[i].second
			p--
		}
		st[p] = ps[i].second
		p++
	}

	odd := make([]bool, n)
	even := make([]bool, n)
	odd[n-1] = true
	even[n-1] = true
	res := 1

	for i := n - 2; i >= 0; i-- {
		if gt[i] > i {
			odd[i] = even[gt[i]]
		}
		if lt[i] > i {
			even[i] = odd[lt[i]]
		}

		if odd[i] {
			res++
		}
	}

	return res

}

type P struct {
	first  int
	second int
}

type PS []P

func (ps PS) Len() int {
	return len(ps)
}

func (ps PS) Less(i, j int) bool {
	return ps[i].first < ps[j].first || (ps[i].first == ps[j].first && ps[i].second < ps[j].second)
}

func (ps PS) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}

type RPS []P

func (ps RPS) Len() int {
	return len(ps)
}

func (ps RPS) Less(i, j int) bool {
	return ps[i].first > ps[j].first || (ps[i].first == ps[j].first && ps[i].second < ps[j].second)
}

func (ps RPS) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}
