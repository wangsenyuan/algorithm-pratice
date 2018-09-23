package p911

import "sort"

type TopVotedCandidate struct {
	ans []Pair
}

func Constructor(persons []int, times []int) TopVotedCandidate {
	n := len(persons)
	ans := make([]Pair, n)
	st := NewSegTree(n + 1)

	for i := 0; i < n; i++ {
		id := persons[i]
		t := times[i]
		top := st.Update(id, 1)
		ans[i] = Pair{t, top}
	}
	return TopVotedCandidate{ans}
}

func (this *TopVotedCandidate) Q(t int) int {
	ans := this.ans
	n := len(ans)
	i := sort.Search(n, func(i int) bool {
		return ans[i].first > t
	})
	i--
	return ans[i].second
}

type Pair struct {
	first, second int
}

type Trip struct {
	first, second, third int
}

type SegTree struct {
	arr  []*Trip
	size int
	time int
}

func NewSegTree(n int) SegTree {
	arr := make([]*Trip, 4*n)

	for i := 0; i < 4*n; i++ {
		arr[i] = new(Trip)
	}

	var loop func(i int, start, end int)
	loop = func(i int, start, end int) {
		if start == end {
			arr[i].third = start
			return
		}
		mid := (start + end) >> 1
		loop(2*i+1, start, mid)
		loop(2*i+2, mid+1, end)
	}

	loop(0, 0, n-1)

	return SegTree{arr, n, 0}
}

func (st *SegTree) Update(pos int, val int) int {
	arr := st.arr
	time := st.time
	var loop func(i int, start, end int)
	loop = func(i int, start, end int) {
		if start == end {
			arr[i].first++
			arr[i].second = time
			return
		}
		mid := (start + end) >> 1
		if pos <= mid {
			loop(2*i+1, start, mid)
		} else {
			loop(2*i+2, mid+1, end)
		}

		if arr[2*i+1].first > arr[2*i+2].first ||
			(arr[2*i+1].first == arr[2*i+2].first && arr[2*i+1].second > arr[2*i+2].second) {
			arr[i].first = arr[2*i+1].first
			arr[i].second = arr[2*i+1].second
			arr[i].third = arr[2*i+1].third
		} else {
			arr[i].first = arr[2*i+2].first
			arr[i].second = arr[2*i+2].second
			arr[i].third = arr[2*i+2].third
		}
	}
	loop(0, 0, st.size-1)
	st.time++
	return arr[0].third
}
