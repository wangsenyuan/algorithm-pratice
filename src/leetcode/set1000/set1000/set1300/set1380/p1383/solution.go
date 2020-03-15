package p1383

import (
	"container/heap"
	"sort"
)

func maxPerformance(n int, speed []int, efficiency []int, k int) int {
	ps := make([]Pair, n)

	for i := 0; i < n; i++ {
		ps[i] = Pair{speed[i], efficiency[i]}
	}

	sort.Sort(Pairs(ps))

	h := make(IntHeap, 0, k)

	heap.Init(&h)

	var res int64

	var total int64
	for i := 0; i < k; i++ {
		total += int64(ps[i].first)

		res = max(res, total*int64(ps[i].second))

		heap.Push(&h, ps[i].first)
	}

	for i := k; i < n; i++ {
		heap.Push(&h, ps[i].first)
		minSpeed := heap.Pop(&h).(int)

		total += int64(ps[i].first) - int64(minSpeed)

		res = max(res, total*int64(ps[i].second))
	}

	return int(res % 1000000007)
}

// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func max(a, b int64) int64 {
	if a >= b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

type Pair struct {
	first, second int
}

type Pairs []Pair

func (this Pairs) Len() int {
	return len(this)
}

func (this Pairs) Less(i, j int) bool {
	return this[i].second > this[j].second
}

func (this Pairs) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}
