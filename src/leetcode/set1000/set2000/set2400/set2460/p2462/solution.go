package p2462

import "container/heap"

func totalCost(costs []int, k int, candidates int) int64 {
	n := len(costs)
	pq := make(PQ, 0, n)

	l, r := 0, n-1
	for i := 0; i < candidates && l <= r; i++ {
		heap.Push(&pq, Employee{costs[i], l})
		l++
		if l <= r {
			heap.Push(&pq, Employee{costs[r], r})
			r--
		}
	}

	var res int64

	for l <= r && k > 0 {
		cur := heap.Pop(&pq).(Employee)
		res += int64(cur.cost)
		if cur.id < l {
			heap.Push(&pq, Employee{costs[l], l})
			l++
		} else {
			heap.Push(&pq, Employee{costs[r], r})
			r--
		}
		k--
	}

	for pq.Len() > 0 && k > 0 {
		res += int64(heap.Pop(&pq).(Employee).cost)
		k--
	}

	return res
}

type Employee struct {
	cost int
	id   int
}

type PQ []Employee

func (pq PQ) Len() int {
	return len(pq)
}

func (pq PQ) Less(i, j int) bool {
	return pq[i].cost < pq[j].cost || pq[i].cost == pq[j].cost && pq[i].id < pq[j].id
}

func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PQ) Push(x interface{}) {
	e := x.(Employee)
	*pq = append(*pq, e)
}

func (pq *PQ) Pop() interface{} {
	old := *pq
	n := len(old)
	ret := old[n-1]
	*pq = old[:n-1]
	return ret
}
