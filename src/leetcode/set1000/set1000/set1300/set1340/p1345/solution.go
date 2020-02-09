package p1345

import "container/heap"

// const INF = 1 << 30

func minJumps(arr []int) int {
	n := len(arr)

	mem := make(map[int][]int)

	for i := 0; i < n; i++ {
		if _, found := mem[arr[i]]; !found {
			mem[arr[i]] = make([]int, 0, 10)
		}
		mem[arr[i]] = append(mem[arr[i]], i)
	}

	items := make([]*Item, n)
	pq := make(PQ, 0, n)

	for i := 0; i < n; i++ {
		items[i] = new(Item)
		items[i].value = i
		items[i].priority = i
		items[i].index = i
		pq = append(pq, items[i])
	}

	heap.Init(&pq)

	for pq.Len() > 0 {
		cur := heap.Pop(&pq).(*Item)
		u := cur.value

		for _, v := range mem[arr[u]] {
			if items[v].index >= 0 && items[v].priority > cur.priority+1 {
				pq.update(items[v], cur.priority+1)
			}
		}

		if u-1 > 0 && items[u-1].index >= 0 && items[u-1].priority > cur.priority+1 {
			pq.update(items[u-1], cur.priority+1)
		}
		if u+1 < n && items[u+1].index >= 0 && items[u+1].priority > cur.priority+1 {
			pq.update(items[u+1], cur.priority+1)
		}
	}

	return items[n-1].priority
}

type Item struct {
	value    int
	priority int
	index    int
}

type PQ []*Item

func (pq PQ) Len() int {
	return len(pq)
}

func (pq PQ) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
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
	n := len(*pq)
	item := (*pq)[n-1]
	(*pq)[n-1] = nil
	item.index = -1

	*pq = (*pq)[:n-1]
	return item
}

func (pq *PQ) update(item *Item, priority int) {
	item.priority = priority
	heap.Fix(pq, item.index)
}
