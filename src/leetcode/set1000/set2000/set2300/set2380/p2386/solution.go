package p2386

import (
	"container/heap"
	"sort"
)

func kSum(nums []int, k int) int64 {
	var sum int64
	for i := 0; i < len(nums); i++ {
		if nums[i] >= 0 {
			sum += int64(nums[i])
		} else {
			// sum -= int64(nums[i])
			nums[i] = -nums[i]
		}
	}

	sort.Ints(nums)

	cell := &Cell{sum: sum, pos: 0}

	pq := make(PriorityQueue, 0, k)

	heap.Push(&pq, cell)

	for k > 1 {
		k--
		cur := heap.Pop(&pq).(*Cell)
		if cur.pos < len(nums) {
			tmp := cur.sum - int64(nums[cur.pos])
			heap.Push(&pq, &Cell{sum: tmp, pos: cur.pos + 1})
			if cur.pos > 0 {
				heap.Push(&pq, &Cell{sum: tmp + int64(nums[cur.pos-1]), pos: cur.pos + 1})
			}
		}
	}
	return pq[0].sum
}

type Cell struct {
	sum int64
	pos int
}

type PriorityQueue []*Cell

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].sum > pq[j].sum
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*Cell)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}
