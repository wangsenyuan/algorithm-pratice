package p1027

import (
	"container/heap"
	"sort"
)

func videoStitching(clips [][]int, T int) int {
	sort.Sort(Pairs(clips))
	//by start time
	time := 0

	pq := new(IntHeap)
	heap.Init(pq)

	var ans int
	n := len(clips)
	var i int
	for time < T {
		for i < n && time >= clips[i][0] {
			heap.Push(pq, clips[i][1])
			i++
		}
		if pq.Len() == 0 {
			break
		}
		time = heap.Pop(pq).(int)
		ans++
	}

	if time < T {
		return -1
	}
	return ans
}

type Pairs [][]int

func (ps Pairs) Len() int {
	return len(ps)
}

func (ps Pairs) Less(i, j int) bool {
	a := ps[i]
	b := ps[j]
	return a[0] < b[0]
}

func (ps Pairs) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
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
