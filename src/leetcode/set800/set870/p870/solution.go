package p870

import "sort"

func advantageCount(A []int, B []int) []int {
	n := len(A)
	bps := make(PS, n)
	for i := 0; i < n; i++ {
		bps[i] = P{B[i], i}
	}
	sort.Sort(bps)
	sort.Ints(A)

	res := make([]int, n)

	for i := 0; i < n; i++ {
		b := bps[i].val
		j := sort.Search(len(A), func(j int) bool {
			return A[j] > b
		})
		if j == len(A) {
			// no one can beat b
			// just take the first one for b
			res[bps[i].idx] = A[0]
			A = A[1:]
		} else {
			res[bps[i].idx] = A[j]
			copy(A[j:], A[j+1:])
			A = A[0 : len(A)-1]
		}
	}

	return res
}

type P struct {
	val int
	idx int
}

type PS []P

func (ps PS) Len() int {
	return len(ps)
}

func (ps PS) Less(i, j int) bool {
	return ps[i].val < ps[j].val
}

func (ps PS) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}
