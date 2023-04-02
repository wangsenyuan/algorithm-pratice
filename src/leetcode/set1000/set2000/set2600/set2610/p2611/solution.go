package p2611

import "container/heap"

func miceAndCheese(reward1 []int, reward2 []int, k int) int {
	n := len(reward1)
	cheese := make([]Pair, n)
	for i := 0; i < n; i++ {
		cheese[i] = Pair{reward1[i], reward2[i]}
	}
	pq := make(PQ, 0, n)
	// k 个被first吃， n - k 个被second吃
	for i := 0; i < n; i++ {
		heap.Push(&pq, &cheese[i])
	}

	var sum int
	for pq.Len() > k {
		x := heap.Pop(&pq).(*Pair)
		sum += x.second
	}

	for pq.Len() > 0 {
		x := heap.Pop(&pq).(*Pair)
		sum += x.first
	}

	return sum
}

type Pair struct {
	first  int
	second int
}

type PQ []*Pair

func (pq PQ) Len() int {
	return len(pq)
}

func (pq PQ) Less(i, j int) bool {
	a := -pq[i].first + pq[i].second
	b := -pq[j].first + pq[j].second
	return a > b
}

func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PQ) Push(x interface{}) {
	it := x.(*Pair)
	*pq = append(*pq, it)
}

func (pq *PQ) Pop() interface{} {
	old := *pq
	n := len(old)
	rt := old[n-1]
	*pq = old[:n-1]
	return rt
}
