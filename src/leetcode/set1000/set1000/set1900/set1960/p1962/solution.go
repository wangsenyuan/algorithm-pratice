package p1962

import "container/heap"

func minStoneSum(piles []int, k int) int {
	pq := make(IntHeap, 0, len(piles))

	for _, cur := range piles {
		heap.Push(&pq, cur)
	}

	for i := 0; i < k; i++ {
		if pq.Len() == 0 {
			break
		}
		cur := heap.Pop(&pq).(int)
		cur -= cur / 2
		if cur > 0 {
			heap.Push(&pq, cur)
		}
	}
	var res int
	for pq.Len() > 0 {
		res += heap.Pop(&pq).(int)
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
