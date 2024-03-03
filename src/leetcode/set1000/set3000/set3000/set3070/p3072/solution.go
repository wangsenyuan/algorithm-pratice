package p3072

import "container/heap"

func minOperations(nums []int, k int) int {
	pq := make(IntHeap, len(nums))
	copy(pq, nums)
	heap.Init(&pq)

	var res int

	for pq[0] < k {
		res++
		x := heap.Pop(&pq).(int)
		y := heap.Pop(&pq).(int)
		z := 2*min(x, y) + max(x, y)
		heap.Push(&pq, z)
	}
	return res
}

type IntHeap []int

func (pq IntHeap) Len() int {
	return len(pq)
}

func (pq IntHeap) Less(i, j int) bool {
	return pq[i] < pq[j]
}

func (pq IntHeap) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *IntHeap) Push(x any) {
	*pq = append(*pq, x.(int))
}

func (pq *IntHeap) Pop() any {
	old := *pq
	n := len(old)
	res := old[n-1]
	*pq = old[:n-1]
	return res
}
