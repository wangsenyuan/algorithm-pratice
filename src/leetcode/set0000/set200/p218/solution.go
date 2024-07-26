package p218

import (
	"container/heap"
	"slices"
)

type pair struct {
	first  int
	second int
	id     int
}

func getSkyline(buildings [][]int) [][]int {
	n := len(buildings)
	items := make([]*Item, n)
	points := make([]pair, 2*n)
	for i, building := range buildings {
		l, r, h := building[0], building[1], building[2]
		points[2*i] = pair{l, -h, i}
		points[2*i+1] = pair{r, h, i}
		it := new(Item)
		it.id = i
		it.priority = 1 << 30
		items[i] = it
	}

	slices.SortFunc(points, func(a, b pair) int {
		if a.first != b.first {
			return a.first - b.first
		}
		return a.second - b.second
	})

	var res [][]int

	pq := make(Heap, 0, n+1)
	var prev int
	heap.Push(&pq, new(Item))

	for _, point := range points {
		x, h, i := point.first, point.second, point.id
		if h < 0 {
			// a left
			items[i].priority = -h
			heap.Push(&pq, items[i])
		} else {
			pq.remove(items[i])
		}

		cur := pq[0].priority

		if cur != prev {
			res = append(res, []int{x, cur})
			prev = cur
		}
	}

	return res
}

type Item struct {
	id       int
	priority int
	index    int
}

type Heap []*Item

func (h Heap) Len() int           { return len(h) }
func (h Heap) Less(i, j int) bool { return h[i].priority > h[j].priority }
func (h Heap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	h[i].index = i
	h[j].index = j
}

func (h *Heap) Push(x any) {
	it := x.(*Item)
	it.index = len(*h)
	*h = append(*h, it)
}

func (h *Heap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *Heap) remove(it *Item) {
	it.priority = 1 << 30
	heap.Fix(h, it.index)
	heap.Pop(h)
}
