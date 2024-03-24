package p3092

import "container/heap"

func mostFrequentIDs(nums []int, freq []int) []int64 {
	items := make(map[int]*Item)

	ans := make([]int64, len(nums))

	pq := make(PriorityQueue, 0, 1)

	for i := 0; i < len(nums); i++ {
		id := nums[i]
		f := freq[i]
		if v, ok := items[id]; ok {
			pq.update(v, v.priority+f)
		} else {
			it := new(Item)
			it.id = id
			it.priority = f
			heap.Push(&pq, it)
			items[id] = it
		}
		if len(pq) > 0 {
			ans[i] = int64(pq[0].priority)
		}
	}

	return ans
}

// An Item is something we manage in a priority queue.
type Item struct {
	id       int
	priority int
	index    int
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority > pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	item := x.(*Item)
	item.index = len(*pq)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue) update(it *Item, v int) {
	it.priority = v
	heap.Fix(pq, it.index)
}
