package main

import (
	"container/heap"
	"fmt"
)

func main() {
	fmt.Println(frequencySort("raaeaedere"))
	fmt.Println(frequencySort("Aabb"))

}

func frequencySort(s string) string {
	frequence := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		frequence[s[i]]++
	}

	pq := make(PQ, 0, len(s))

	for b, f := range frequence {
		pq = append(pq, &Word{b, f})
	}

	heap.Init(&pq)
	var res []byte
	for len(pq) > 0 {
		v := heap.Pop(&pq).(*Word)
		for i := 0; i < v.f; i++ {
			res = append(res, v.b)
		}
	}
	return string(res)
}

type Word struct {
	b byte
	f int
}

type PQ []*Word

func (pq PQ) Len() int {
	return len(pq)
}

func (pq PQ) Less(i, j int) bool {
	return pq[i].f > pq[j].f
}

func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PQ) Push(x interface{}) {
	*pq = append(*pq, x.(*Word))
}

func (pq *PQ) Pop() interface{} {
	old := *pq
	n := len(old)
	word := old[n-1]
	*pq = old[0 : n-1]
	return word
}
