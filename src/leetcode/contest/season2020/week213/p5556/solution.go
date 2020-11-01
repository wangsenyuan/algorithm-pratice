package p5556

import "container/heap"

func furthestBuilding(heights []int, bricks int, ladders int) int {
	n := len(heights)

	pq := make(IntHeap, 0, n-1)
	var i int
	for ; i < n-1; i++ {
		if heights[i] >= heights[i+1] {
			continue
		}
		d := heights[i+1] - heights[i]
		heap.Push(&pq, d)
		bricks -= d
		if bricks < 0 {
			if ladders == 0 {
				break
			}
			// try to exchange bricks with ladder
			x := heap.Pop(&pq).(int)
			bricks += x
			ladders--
		}
	}

	return i
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
