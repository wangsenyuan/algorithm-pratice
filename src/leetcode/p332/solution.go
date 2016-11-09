package main

import (
	"container/heap"
	"fmt"
)

func main() {
	tickets := [][]string{[]string{"JFK", "SFO"}, []string{"JFK", "ATL"}, []string{"SFO", "ATL"}, []string{"ATL", "JFK"}, []string{"ATL", "SFO"}}
	fmt.Println(findItinerary(tickets))
}

func findItinerary(tickets [][]string) []string {
	targets := make(map[string]*StrPQ)

	for _, ticket := range tickets {
		from, to := ticket[0], ticket[1]
		if targets[from] == nil {
			targets[from] = &StrPQ{}
			heap.Init(targets[from])
		}
		heap.Push(targets[from], to)
	}

	var ans []string
	var visit func(from string)

	visit = func(from string) {
		if pq, ok := targets[from]; ok {
			for pq.Len() > 0 {
				visit(heap.Pop(pq).(string))
			}
			ans = append(ans, from)
		} else {
			ans = append(ans, from)
		}
	}
	visit("JFK")

	return reverse(ans)
}

func reverse(strs []string) []string {
	res := make([]string, len(strs))
	for i, str := range strs {
		res[len(strs)-i-1] = str
	}
	return res
}

type StrPQ []string

func (h StrPQ) Len() int           { return len(h) }
func (h StrPQ) Less(i, j int) bool { return h[i] < h[j] }
func (h StrPQ) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *StrPQ) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(string))
}

func (h *StrPQ) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
