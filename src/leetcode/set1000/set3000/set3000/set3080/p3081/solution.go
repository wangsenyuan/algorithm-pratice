package p3081

import "container/heap"

func unmarkedSumArray(nums []int, queries [][]int) []int64 {
	n := len(nums)
	mark := make([]*Pair, n)
	pq := make(Heap, n)
	var sum int64
	for i := 0; i < n; i++ {
		it := &Pair{nums[i], i, i}
		mark[i] = it
		pq[i] = it
		sum += int64(nums[i])
	}

	heap.Init(&pq)

	ans := make([]int64, len(queries))

	for i, cur := range queries {
		j, k := cur[0], cur[1]
		if mark[j].index >= 0 {
			sum -= int64(mark[j].first)
			pq.update(mark[j], 0)
			heap.Pop(&pq)
		}
		for k > 0 && pq.Len() > 0 {
			k--
			sum -= int64(heap.Pop(&pq).(*Pair).first)
		}
		ans[i] = sum
	}

	return ans
}

type Pair struct {
	first  int
	second int
	index  int
}

type Heap []*Pair

func (pq Heap) Len() int {
	return len(pq)
}

func (pq Heap) Less(i, j int) bool {
	return pq[i].first < pq[j].first || pq[i].first == pq[j].first && pq[i].second < pq[j].second
}

func (pq Heap) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *Heap) Push(x any) {
	p := x.(*Pair)
	p.index = len(*pq)
	*pq = append(*pq, p)
}

func (pq *Heap) Pop() any {
	old := *pq
	n := len(old)
	res := old[n-1]
	res.index = -1
	*pq = old[:n-1]
	return res
}

func (pq *Heap) update(x *Pair, v int) {
	x.first = v
	heap.Fix(pq, x.index)
}
