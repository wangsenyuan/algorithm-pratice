package p3362

import "container/heap"

func maxRemoval(nums []int, queries [][]int) int {
	n := len(nums)

	open := make([][]int, n)
	end := make([][]int, n)

	for i, cur := range queries {
		l, r := cur[0], cur[1]
		open[l] = append(open[l], i)
		end[r] = append(end[r], i)
	}

	items := make([]*Item, len(queries))

	active := make(PriorityQueue, 0, len(queries))
	var res int

	add := make([]int, n+1)

	for pos := 0; pos < n; pos++ {
		for _, i := range open[pos] {
			r := queries[i][1]
			it := new(Item)
			it.id = i
			it.priority = r
			heap.Push(&active, it)
			items[i] = it
		}

		if pos > 0 {
			add[pos] += add[pos-1]
		}

		v := nums[pos]

		for v > add[pos] {
			if len(active) == 0 {
				return -1
			}
			add[pos]++
			it := heap.Pop(&active).(*Item)
			add[it.priority+1]--
		}
		// 当前位置结束的必须处理掉
		for _, i := range end[pos] {
			if items[i].index >= 0 {
				res++
				active.update(items[i], 1<<30)
				heap.Pop(&active)
			}
		}
	}
	// 剩余的没有使用的，都可以清理掉
	return res + len(active)
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

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	item.index = -1
	return item
}

func (pq *PriorityQueue) update(it *Item, v int) {
	it.priority = v
	heap.Fix(pq, it.index)
}
