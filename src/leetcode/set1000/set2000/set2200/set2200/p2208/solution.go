package p2208

import "container/heap"

func halveArray(nums []int) int {
	n := len(nums)
	pq := make(PQ, n)

	var sum float64

	for i := 0; i < n; i++ {
		item := new(Item)
		item.num = int64(nums[i])
		item.div = 1
		item.index = i
		sum += float64(nums[i])
		pq[i] = item
	}

	heap.Init(&pq)

	var sub float64

	var cnt int
	for 2*sub < sum {
		top := heap.Pop(&pq).(*Item)
		top.div *= 2
		sub += float64(top.num) / float64(top.div)
		heap.Push(&pq, top)
		cnt++
	}

	return cnt
}

type Item struct {
	num   int64
	div   int64
	index int
}

func (this Item) Less(that Item) bool {
	// num / div
	return this.num*that.div > this.div*that.num
}

type PQ []*Item

func (pq PQ) Len() int {
	return len(pq)
}

func (pq PQ) Less(i, j int) bool {
	return pq[i].Less(*pq[j])
}

func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PQ) Push(x interface{}) {
	item := x.(*Item)
	item.index = len(*pq)
	*pq = append(*pq, item)
}

func (pq *PQ) Pop() interface{} {
	old := *pq
	n := len(old)
	ret := old[n-1]
	ret.index = -1
	*pq = old[:n-1]
	return ret
}
