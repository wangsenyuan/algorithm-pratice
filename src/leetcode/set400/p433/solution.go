package main

import (
	"container/heap"
	"fmt"
)

func main() {
	bank := []string{"AACCGGTA", "AACCGCTA", "AAACGGTA"}
	start := "AACCGGTT"
	end := "AAACGGTA"
	fmt.Println(minMutation(start, end, bank))
}

var INF = 10000000

func minMutation(start string, end string, bank []string) int {
	pq := make(PriorityQueue, len(bank))
	itemMap := make(map[string]*Item)
	for i, gene := range bank {
		pq[i] = &Item{gene: gene, priority: INF, index: i}
		itemMap[gene] = pq[i]
	}

	if _, canEnd := itemMap[end]; !canEnd {
		return -1
	}

	heap.Init(&pq)
	pq.update(itemMap[end], 0)

	checked := make(map[string]bool)

	for pq.Len() > 0 {
		v := heap.Pop(&pq).(*Item)
		checked[v.gene] = true
		for _, gene := range bank {
			if checked[gene] {
				continue
			}
			w := itemMap[gene]
			if differOne(v.gene, gene) && w.priority > v.priority+1 {
				pq.update(w, v.priority+1)
			}
		}
	}

	ans := -1
	for _, gene := range bank {
		v := itemMap[gene]
		if differOne(gene, start) && (ans == -1 || ans > v.priority+1) {
			ans = v.priority + 1
		}
	}

	return ans
}

func differOne(a, b string) bool {
	x := 0

	for i := range a {
		if a[i] != b[i] {
			x++
		}
	}

	return x == 1
}

type Item struct {
	gene     string
	priority int
	index    int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
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
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) update(item *Item, priority int) {
	item.priority = priority
	heap.Fix(pq, item.index)
}
