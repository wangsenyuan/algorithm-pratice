package p3169

import "container/heap"

func clearStars(s string) string {
	n := len(s)
	pq := make(PQ, 0, n)
	del := make([]bool, n)

	for i := 0; i < n; i++ {
		if s[i] == '*' {
			del[i] = true
			it := heap.Pop(&pq).(*Item)
			del[it.pos] = true
		} else {
			it := new(Item)
			it.pos = i
			it.val = int(s[i] - 'a')
			heap.Push(&pq, it)
		}
	}

	buf := []byte(s)
	var j int

	for i := 0; i < n; i++ {
		if !del[i] {
			buf[j] = s[i]
			j++
		}
	}

	return string(buf[:j])
}

type Item struct {
	pos int
	val int
}

type PQ []*Item

func (pq PQ) Len() int {
	return len(pq)
}

func (pq PQ) Less(i, j int) bool {
	return pq[i].val < pq[j].val || pq[i].val == pq[j].val && pq[i].pos > pq[j].pos
}

func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PQ) Push(x any) {
	it := x.(*Item)
	*pq = append(*pq, it)
}

func (pq *PQ) Pop() any {
	old := *pq
	n := len(old)
	ret := old[n-1]
	*pq = old[:n-1]
	return ret
}
