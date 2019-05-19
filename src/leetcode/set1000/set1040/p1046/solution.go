package p1046

import "container/heap"

func lastStoneWeight(stones []int) int {
	que := IntHeap(stones)
	heap.Init(&que)

	for len(que) > 1 {
		first := heap.Pop(&que).(int)
		second := heap.Pop(&que).(int)
		if first == second {
			continue
		}
		heap.Push(&que, abs(first - second))
	}
	if que.Len() == 1 {
		return heap.Pop(&que).(int)
	}
	return 0
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

type IntHeap []int

func (this IntHeap) Len() int {
	return len(this)
}

func (this IntHeap) Less(i, j int) bool {
	return this[i] > this[j]
}

func (this IntHeap) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func (this *IntHeap) Push(item interface{}) {
	*this = append(*this, item.(int))
}

func (this *IntHeap) Pop() interface{} {
	n := len(*this)
	last := (*this)[n-1]
	*this = (*this)[:n-1]
	return last
}