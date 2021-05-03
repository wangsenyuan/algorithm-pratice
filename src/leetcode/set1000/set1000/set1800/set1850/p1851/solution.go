package p1851

import (
	"container/heap"
	"sort"
)

func minInterval(intervals [][]int, queries []int) []int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	qs := make([]Pair, len(queries))
	for i, q := range queries {
		qs[i] = Pair{q, i}
	}

	sort.Slice(qs, func(i, j int) bool {
		return qs[i].first < qs[j].first
	})

	ans := make([]int, len(qs))

	right := make(PQ, 0, len(intervals))
	active := make(PQ, 0, len(intervals))
	items := make([]*Item, len(intervals))
	var j int
	for _, q := range qs {
		for right.Len() > 0 {
			tmp := right[0].key
			if tmp.first >= q.first {
				break
			}
			r := tmp.second
			active.remove(items[r])
			heap.Pop(&right)
		}

		for j < len(intervals) && intervals[j][0] <= q.first {
			if intervals[j][1] >= q.first {
				items[j] = new(Item)
				items[j].key = Pair{intervals[j][1] - intervals[j][0] + 1, j}
				heap.Push(&active, items[j])
				ritem := new(Item)
				ritem.key = Pair{intervals[j][1], j}
				heap.Push(&right, ritem)
			}
			j++
		}

		if active.Len() == 0 {
			ans[q.second] = -1
		} else {
			ans[q.second] = active[0].key.first
		}
	}
	return ans
}

type Pair struct {
	first, second int
}

func (this Pair) Less(that Pair) bool {
	return this.first < that.first || this.first == that.first && this.second < that.second
}

type Item struct {
	key   Pair
	index int
}

type PQ []*Item

func (pq PQ) Len() int {
	return len(pq)
}

func (pq PQ) Less(i, j int) bool {
	return pq[i].key.Less(pq[j].key)
}

func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PQ) Push(x interface{}) {
	cur := x.(*Item)
	cur.index = len(*pq)
	*pq = append(*pq, cur)
}

func (pq *PQ) Pop() interface{} {
	old := *pq
	n := len(old)
	res := old[n-1]
	res.index = -1
	*pq = old[:n-1]
	return res
}

func (pq *PQ) update(item *Item, v Pair) {
	item.key = v
	heap.Fix(pq, item.index)
}

func (pq *PQ) remove(item *Item) {
	pq.update(item, Pair{0, -1})
	heap.Pop(pq)
}
