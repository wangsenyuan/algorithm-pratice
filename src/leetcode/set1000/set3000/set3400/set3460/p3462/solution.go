package p3462

import (
	"container/heap"
	"sort"
)

func maxSum(grid [][]int, limits []int, k int) int64 {
	pq := make(IntHeap, 0, k+1)

	for i, row := range grid {
		sort.Ints(row)
		for j := len(row) - 1; j >= 0 && limits[i] > 0; j-- {
			limits[i]--
			heap.Push(&pq, row[j])
		}
	}
	var res int

	for pq.Len() > 0 && k > 0 {
		res += heap.Pop(&pq).(int)
		k--
	}
	return int64(res)
}

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
