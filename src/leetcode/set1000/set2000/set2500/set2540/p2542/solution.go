package p2542

import (
	"container/heap"
	"sort"
)

func maxScore(A []int, B []int, k int) int64 {
	n := len(A)
	// n ~ 1e5
	P := make([]Pair, n)
	for i := 0; i < n; i++ {
		P[i] = Pair{A[i], B[i]}
	}
	sort.Slice(P, func(i, j int) bool {
		return P[i].second > P[j].second
	})

	pq := make(IntHeap, 0, k)
	var sum int64
	for i := 0; i < k-1; i++ {
		heap.Push(&pq, P[i].first)
		sum += int64(P[i].first)
	}

	var res int64

	for i := k - 1; i < n; i++ {
		heap.Push(&pq, P[i].first)
		sum += int64(P[i].first)
		res = max(res, sum*int64(P[i].second))
		min := heap.Pop(&pq).(int)
		sum -= int64(min)
	}

	return res
}

func max(a, b int64) int64 {
	if a >= b {
		return a
	}
	return b
}

type Pair struct {
	first  int
	second int
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
