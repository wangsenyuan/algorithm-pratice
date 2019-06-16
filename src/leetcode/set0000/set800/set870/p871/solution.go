package p871

import "container/heap"

func minRefuelStops(target int, startFuel int, stations [][]int) int {
	n := len(stations)
	hp := make(IntHeap, 0, n)
	fuel := startFuel
	var reach int
	for i := 0; i < n; i++ {
		need := stations[i][0]
		if i > 0 {
			need -= stations[i-1][0]
		}
		for hp.Len() > 0 && fuel < need {
			fuel += heap.Pop(&hp).(int)
		}
		if fuel < need {
			return -1
		}
		fuel -= need
		reach = stations[i][0]
		if reach >= target {
			break
		}
		// try not to use this stations fuel
		heap.Push(&hp, stations[i][1])
	}

	for hp.Len() > 0 && reach+fuel < target {
		fuel += heap.Pop(&hp).(int)
	}

	if reach+fuel < target {
		return -1
	}

	return n - hp.Len()
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
