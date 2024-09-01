package p3275

import "container/heap"

func resultsArray(queries [][]int, k int) []int {
	res := make([]int, len(queries))
	pq := make(IntHeap, 0, len(queries))

	for i, cur := range queries {
		x := abs(cur[0])
		y := abs(cur[1])

		if pq.Len() < k || x+y <= pq[0] {
			heap.Push(&pq, x+y)
			if pq.Len() == k+1 {
				heap.Pop(&pq)
			}
		}

		res[i] = -1

		if len(pq) == k {
			res[i] = pq[0]
		}
	}

	return res
}

func abs(num int) int {
	return max(num, -num)
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
