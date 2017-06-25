package main

import (
	"sort"
	"fmt"
	"container/heap"
)

func scheduleCourse(courses [][]int) int {

	sort.Sort(ByEnd(courses))

	h := &IntHeap{}
	heap.Init(h)

	cur := 0

	for _, course := range courses {
		if cur+course[0] <= course[1] {
			heap.Push(h, course[0])
			cur += course[0]
		} else if h.Len() > 0 {
			longest := heap.Pop(h).(int)
			if cur-longest+course[0] <= course[1] && longest > course[0] {
				heap.Push(h, course[0])
				cur = cur - longest + course[0]
			} else {
				heap.Push(h, longest)
			}
		}
	}

	return h.Len()
}

type ByEnd [][]int

func (this ByEnd) Len() int {
	return len(this)
}

func (this ByEnd) Less(i, j int) bool {
	return this[i][1] < this[j][1]
}

func (this ByEnd) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
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
	*h = old[0: n-1]
	return x
}

func main() {
	//courses := [][]int{{100, 200}, {200, 1300}, {1000, 1250}, {2000, 3200}}
	//courses := [][]int{{5, 5}, {4, 6}, {2, 6}}
	//courses := [][]int{{5, 15}, {3, 19}, {6, 7}, {2, 10}, {5, 16}, {8, 14}, {10, 11}, {2, 19}}
	courses := [][]int{{7, 17}, {3, 12}, {10, 20}, {9, 10}, {5, 20}, {10, 19}, {4, 18}}

	fmt.Println(scheduleCourse(courses))
}
