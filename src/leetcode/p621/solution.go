package main

import (
	"container/heap"
	"fmt"
)

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

func leastInterval(tasks []byte, n int) int {
	h := &IntHeap{}
	heap.Init(h)

	cnts := make([]int, 26)
	for _, task := range tasks {
		cnts[task-'A']++
	}

	for _, cnt := range cnts {
		if cnt > 0 {
			heap.Push(h, cnt)
		}
	}

	var time int
	for h.Len() > 0 {
		var tmp []int
		for i := 0; i <= n; i++ {
			if h.Len() > 0 {
				cnt := heap.Pop(h).(int)
				if cnt > 1 {
					tmp = append(tmp, cnt-1)
				}
			}
			time++

			if h.Len() == 0 && len(tmp) == 0 {
				// no more iteration need
				break
			}
		}

		for _, x := range tmp {
			heap.Push(h, x)
		}
	}

	return time
}

func main() {
	tasks := []byte{'A', 'A', 'A', 'B', 'B', 'B'}
	n := 2
	fmt.Println(leastInterval(tasks, n))
}
