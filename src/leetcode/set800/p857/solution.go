package p857

import (
	"container/heap"
	"math"
	"sort"
)

func mincostToHireWorkers(quality []int, wage []int, K int) float64 {
	n := len(quality)
	workers := make(Workers, n)

	for i := 0; i < n; i++ {
		workers[i] = Worker{wage[i], quality[i], float64(wage[i]) / float64(quality[i])}
	}

	sort.Sort(workers)

	hp := make(IntHeap, 0, n)
	var sum int64
	var res = math.MaxFloat64
	for _, worker := range workers {
		q, r := worker.quality, worker.ratio
		sum += int64(q)
		heap.Push(&hp, q)
		if hp.Len() > K {
			top := heap.Pop(&hp).(int)
			sum -= int64(top)
		}
		if hp.Len() == K {
			res = math.Min(res, float64(sum)*r)
		}
	}
	return res
}

type Worker struct {
	wage    int
	quality int
	ratio   float64
}

type Workers []Worker

func (ws Workers) Len() int {
	return len(ws)
}

func (ws Workers) Less(i, j int) bool {
	return ws[i].ratio < ws[j].ratio
}

func (ws Workers) Swap(i, j int) {
	ws[i], ws[j] = ws[j], ws[i]
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
