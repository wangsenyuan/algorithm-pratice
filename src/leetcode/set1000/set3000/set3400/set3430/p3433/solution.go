package p3433

import (
	"container/heap"
	"strconv"
	"strings"
)

func countMentions(n int, events [][]string) []int {
	res := make([]int, n)

	pq := make(PriorityQueue, 0, n)
	for i, event := range events {
		time, _ := strconv.Atoi(event[1])
		//ids := event[2]
		it := new(Item)
		it.id = i
		it.event = event[0]
		it.priority = time
		heap.Push(&pq, it)
	}

	var all int

	online := make(map[int]bool)

	for i := range n {
		online[i] = true
	}

	backAt := make([]int, n)
	for i := 0; i < n; i++ {
		backAt[i] = -1
	}

	for pq.Len() > 0 {
		it := heap.Pop(&pq).(*Item)
		i := it.id
		if i >= 0 {
			event := events[i]
			if event[0] == "MESSAGE" {
				if event[2] == "ALL" {
					all++
				} else if event[2] == "HERE" {
					for id := range online {
						res[id]++
					}
				} else {
					ss := strings.Split(event[2], " ")
					for _, s := range ss {
						// idx
						id, _ := strconv.Atoi(s[2:])
						res[id]++
					}
				}
			} else {
				// OFFLINE
				id, _ := strconv.Atoi(event[2])
				delete(online, id)
				nit := new(Item)
				nit.id = -(id + 1)
				nit.priority = it.priority + 60
				backAt[id] = nit.priority
				heap.Push(&pq, nit)
			}
		} else {
			// an online operation
			id := -i - 1
			if backAt[id] == it.priority {
				online[id] = true
			}
		}
	}

	for i := range n {
		res[i] += all
	}
	return res
}

// An Item is something we manage in a priority queue.
type Item struct {
	id       int // The value of the item; arbitrary.
	event    string
	priority int // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority ||
		pq[i].priority == pq[j].priority && pq[i].event != "MESSAGE"
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}
