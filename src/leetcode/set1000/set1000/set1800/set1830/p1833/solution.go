package p1833

import (
	"container/heap"
	"sort"
)

func getOrder(tasks [][]int) []int {
	tt := make([]Task, len(tasks))

	for i, task := range tasks {
		tt[i] = Task{task[0], task[1], i}
	}

	sort.Slice(tt, func(i, j int) bool {
		return tt[i].start < tt[j].start
	})

	que := make(PriorityQueue, 0, len(tasks))
	res := make([]int, 0, len(tt))
	var cur int
	var i int
	for {
		if que.Len() == 0 && i < len(tt) && tt[i].start > cur {
			cur = tt[i].start
		}
		for i < len(tt) && tt[i].start <= cur {
			job := new(Item)
			job.value = tt[i].id
			job.priority = tt[i].process
			heap.Push(&que, job)
			i++
		}
		if que.Len() == 0 {
			break
		}
		first := heap.Pop(&que).(*Item)
		id := first.value
		res = append(res, id)
		cur += tasks[id][1]
	}

	return res
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

type Task struct {
	start, process int
	id             int
}

// An Item is something we manage in a priority queue.
type Item struct {
	value    int // The value of the item; arbitrary.
	priority int // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority || pq[i].priority == pq[j].priority && pq[i].value < pq[j].value
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}
