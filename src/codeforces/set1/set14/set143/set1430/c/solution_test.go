package main

import (
	"container/heap"
	"testing"
)

func runSample(t *testing.T, n int, expect int) {
	ans, res := solve(n)
	if ans != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, ans)
	}
	if len(res) != n-1 {
		t.Fatalf("Sample result is length %d, not %d - 1", len(res), n)
	}

	pq := make(IntHeap, 0, n)
	for i := 0; i < n; i++ {
		heap.Push(&pq, i+1)
	}

	pop := func(num int) bool {
		var arr []int
		for pq.Len() > 0 {
			x := heap.Pop(&pq).(int)
			if x == num {
				for _, y := range arr {
					heap.Push(&pq, y)
				}
				return true
			}
			arr = append(arr, x)
		}
		return false
	}

	for _, cur := range res {
		a, b := cur[0], cur[1]
		if !pop(a) {
			t.Fatalf("no num %d found to add", a)
		}
		if !pop(b) {
			t.Fatalf("no num %d found to add", b)
		}
		heap.Push(&pq, (a+b+1)/2)
	}

	if pq.Len() != 1 || pq[0] != ans {
		t.Fatalf("Sample got not wrong result ")
	}
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

func TestSample1(t *testing.T) {
	runSample(t, 4, 2)
}

func TestSample2(t *testing.T) {
	runSample(t, 5, 2)
}

func TestSample3(t *testing.T) {
	runSample(t, 6, 2)
}
