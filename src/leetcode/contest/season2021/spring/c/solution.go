package c

import "container/heap"

func magicTower(nums []int) int {
	n := len(nums)

	var cur = 1

	pq := make(IntHeap, 0, n)

	bad := make([]int, 0, n)
	var res int
	for i := 0; i < n; i++ {
		if nums[i] < 0 {
			heap.Push(&pq, nums[i])
		}
		cur += nums[i]
		for cur < 0 && pq.Len() > 0 {
			// need to move it to end
			most := heap.Pop(&pq).(int)
			cur -= most
			bad = append(bad, most)
			res++
		}
		if cur < 0 {
			return -1
		}
	}

	for i := 0; i < len(bad); i++ {
		cur += bad[i]
		if cur < 0 {
			return -1
		}
	}
	return res
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
