package p2563

import (
	"container/heap"
	"math"
)

func pickGifts(gifts []int, k int) int64 {
	pq := make(IntHeap, len(gifts))
	copy(pq, gifts)
	heap.Init(&pq)
	for i := 0; i < k && pq.Len() > 0; i++ {
		x := heap.Pop(&pq).(int)
		x = int(math.Sqrt(float64(x)))
		if x > 0 {
			heap.Push(&pq, x)
		}
	}

	var res int64
	for pq.Len() > 0 {
		res += int64(heap.Pop(&pq).(int))
	}

	return res
}

// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
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
