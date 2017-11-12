package main

import (
	"container/heap"
	"fmt"
)

func findOrder(numCourses int, prerequisites [][]int) []int {
	graph := make([]map[int]bool, numCourses)
	rgraph := make([]map[int]bool, numCourses)
	for _, pre := range prerequisites {
		a, b := pre[0], pre[1]
		if graph[b] == nil {
			graph[b] = make(map[int]bool)
		}
		graph[b][a] = true

		if rgraph[a] == nil {
			rgraph[a] = make(map[int]bool)
		}
		rgraph[a][b] = true
	}

	checked := make([]bool, numCourses)

	res := make([]int, 0, numCourses)

	pq := make(PQ, numCourses)
	items := make(map[int]*Item)

	for i := 0; i < numCourses; i++ {
		pq[i] = &Item{at: i, value: rgraph[i], index: i}
		items[i] = pq[i]
	}

	heap.Init(&pq)

	for len(res) < numCourses && pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)

		if len(item.value) > 0 {
			break
		}
		v := item.at
		if v == -1 {
			break
		}
		res = append(res, v)
		checked[v] = true
		for w := range graph[v] {
			delete(graph[v], w)
			delete(rgraph[w], v)
			pq.update(items[w], rgraph[w])
		}
	}

	if len(res) < numCourses {
		return []int{}
	}
	return res
}

type Item struct {
	at    int
	value map[int]bool
	index int
}

type PQ []*Item

func (pq PQ) Len() int {
	return len(pq)
}

func (pq PQ) Less(i, j int) bool {
	return len(pq[i].value) < len(pq[j].value)
}

func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PQ) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PQ) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (pq *PQ) update(item *Item, value map[int]bool) {
	item.value = value
	heap.Fix(pq, item.index)
}

func main() {
	pres := [][]int{[]int{1, 0}, []int{2, 0}, []int{3, 1}, []int{3, 2}}
	//pres := [][]int{[]int{1, 0}, {0, 1}}
	res := findOrder(4, pres)
	fmt.Println(res)
}
