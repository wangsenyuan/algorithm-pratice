package lcp68

import "container/heap"

const MOD = 1e9 + 7

func modAdd(a, b int) int {
	a += b
	if a >= MOD {
		a -= MOD
	}
	return a
}

const X = 100010

func beautifulBouquet(flowers []int, cnt int) int {
	// 以i结尾，长度为l满足条件时，长度为l-1也满足条件
	// 对于任意的i，需要找到满足条件的最左边的l
	// 如果L[i]表示i的l，那么L[i+1] >= L[i]
	var res int
	its := make([]*Item, X)
	pq := make(PQ, X)

	for i := 0; i < X; i++ {
		its[i] = new(Item)
		its[i].id = i
		its[i].priority = 0
		its[i].index = i
		pq[i] = its[i]
	}
	n := len(flowers)
	for r, l := 0, 0; r < n; r++ {
		cur := its[flowers[r]].priority + 1
		pq.Update(its[flowers[r]], cur)
		for l < r && pq[0].priority > cnt {
			tmp := its[flowers[l]].priority - 1
			pq.Update(its[flowers[l]], tmp)
			l++
		}
		res = modAdd(res, r-l+1)
	}

	return res
}

type Item struct {
	id       int
	priority int
	index    int
}

func (this Item) Less(that Item) bool {
	return this.priority > that.priority
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
	i := x.(*Item)
	i.index = len(*pq)
	*pq = append(*pq, i)
}

func (pq *PQ) Pop() interface{} {
	old := *pq
	n := len(old)
	ret := old[n-1]
	*pq = old[:n-1]
	return ret
}

func (pq *PQ) Update(i *Item, p int) {
	i.priority = p
	heap.Fix(pq, i.index)
}
