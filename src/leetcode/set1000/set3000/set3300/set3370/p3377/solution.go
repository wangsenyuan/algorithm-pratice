package p3377

import (
	"container/heap"
	"strconv"
)

func minOperations(n int, m int) int {
	x := max(n, m) * 10
	// 完全没有想法
	lps := make([]int, x+1)
	var primes []int
	for i := 2; i <= x; i++ {
		if lps[i] == 0 {
			primes = append(primes, i)
			lps[i] = i
		}
		for _, j := range primes {
			if i*j > x {
				break
			}
			lps[i*j] = j
			if i%j == 0 {
				break
			}
		}
	}
	if lps[m] == m || lps[n] == n {
		return -1
	}

	if n == m {
		return n
	}

	if m == 1 {
		// 肯定要经过2
		return -1
	}

	items := make(map[int]*Item)
	pq := make(PriorityQueue, 0, 10)
	items[n] = NewItem(n, n)
	heap.Push(&pq, items[n])

	update := func(s string, p int) {
		v, _ := strconv.Atoi(s)
		if lps[v] == v {
			return
		}
		p += v
		if it, ok := items[v]; !ok {
			items[v] = NewItem(v, p)
			heap.Push(&pq, items[v])
		} else if it.priority > p {
			pq.update(it, p)
		}
	}

	for pq.Len() > 0 {
		it := heap.Pop(&pq).(*Item)
		if it.id == m {
			return it.priority
		}
		u := it.id
		// change nums
		ds := []byte(strconv.Itoa(u))

		for i := 0; i < len(ds); i++ {
			if ds[i] != '9' {
				ds[i]++
				update(string(ds), it.priority)
				ds[i]--
			}
			if ds[i] != '0' {
				ds[i]--
				update(string(ds), it.priority)
				ds[i]++
			}
		}
	}

	return -1
}

// An Item is something we manage in a priority queue.
type Item struct {
	id       int
	priority int
	index    int
}

func NewItem(id int, p int) *Item {
	it := new(Item)
	it.id = id
	it.priority = p
	return it
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	item := x.(*Item)
	item.index = len(*pq)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	item.index = -1
	return item
}

func (pq *PriorityQueue) update(it *Item, v int) {
	it.priority = v
	heap.Fix(pq, it.index)
}
