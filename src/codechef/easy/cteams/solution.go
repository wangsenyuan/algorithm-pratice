package main

import (
	"container/heap"
	"fmt"
)

func main() {
	var n int
	fmt.Scanf("%d", &n)

	chefs := make([]Chef, n)
	for i := 0; i < n; i++ {
		var age int
		var rating int64
		fmt.Scanf("%d %d", &age, &rating)

		chefs[i] = Chef{age, rating}
	}

	res := solve(n, chefs)

	for i := 0; i < n; i++ {
		fmt.Println(res[i])
	}
}

func solve(n int, input []Chef) []int64 {
	result := make([]int64, n)
	young, old := int64(0), int64(0)

	youngHeap := make(YoungHeap, 0, n/2)
	oldHeap := make(OldHeap, 0, n/2)

	heap.Init(&youngHeap)
	heap.Init(&oldHeap)

	for i := 0; i < n; i++ {
		age, rating := input[i].age, input[i].rating

		if youngHeap.Len() == 0 {
			heap.Push(&youngHeap, input[i])
			young += rating
		} else {
			tmp := youngHeap[len(youngHeap)-1]

			if tmp.age < age {
				// put current into old
				heap.Push(&oldHeap, input[i])
				old += rating
			} else {
				heap.Push(&youngHeap, input[i])
				young += rating
			}
		}

		for youngHeap.Len() > oldHeap.Len()+1 {
			tmp := heap.Pop(&youngHeap).(Chef)
			young -= tmp.rating
			heap.Push(&oldHeap, tmp)
			old += tmp.rating
		}

		for oldHeap.Len() > youngHeap.Len() {
			tmp := heap.Pop(&oldHeap).(Chef)
			old -= tmp.rating
			heap.Push(&youngHeap, tmp)
			young += tmp.rating
		}

		result[i] = abs(old - young)
	}

	return result
}

func abs(x int64) int64 {
	if x >= 0 {
		return x
	}
	return -x
}

//Chef ..
type Chef struct {
	age    int
	rating int64
}

type OldHeap []Chef

func (h OldHeap) Len() int           { return len(h) }
func (h OldHeap) Less(i, j int) bool { return h[i].age < h[j].age }
func (h OldHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *OldHeap) Push(x interface{}) {
	*h = append(*h, x.(Chef))
}

func (h *OldHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type YoungHeap []Chef

func (h YoungHeap) Len() int           { return len(h) }
func (h YoungHeap) Less(i, j int) bool { return h[i].age > h[j].age }
func (h YoungHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *YoungHeap) Push(x interface{}) {
	*h = append(*h, x.(Chef))
}

func (h *YoungHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
