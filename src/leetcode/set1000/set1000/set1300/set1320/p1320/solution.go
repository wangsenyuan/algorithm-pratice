package p1320

import "container/heap"

const INF = 1 << 20

func minimumDistance(word string) int {
	keyword := []string{
		"ABCDEF", "GHIJKL", "MNOPQR", "STUVWX", "YZ",
	}

	pos := make([][]int, 26)

	for i := 0; i < len(keyword); i++ {
		for j := 0; j < len(keyword[i]); j++ {
			x := int(keyword[i][j] - 'A')
			pos[x] = []int{i, j}
		}
	}

	getDist := func(i, j int) int {
		x := int(word[i] - 'A')
		y := int(word[j] - 'A')
		return abs(pos[x][0]-pos[y][0]) + abs(pos[x][1]-pos[y][1])
	}

	n := len(word)
	items := make([][]*Item, n+1)

	pq := make(PriorityQueue, (n+1)*(n+1))

	for i := 0; i <= n; i++ {
		items[i] = make([]*Item, n+1)
		for j := 0; j <= n; j++ {
			item := new(Item)
			item.priority = INF
			item.value = i*(n+1) + j
			items[i][j] = item
			item.index = item.value
			pq[item.value] = item
		}
	}
	items[0][0].priority = 0

	heap.Init(&pq)

	for pq.Len() > 0 {
		cur := heap.Pop(&pq).(*Item)

		x, y := cur.value/(n+1), cur.value%(n+1)
		d := cur.priority

		if x == n || y == n {
			return d
		}

		nxt := max(x, y) + 1

		if x > 0 {
			dd := getDist(x-1, nxt-1) + d
			if items[nxt][y].priority > dd {
				pq.update(items[nxt][y], dd)
			}
		} else {
			if items[nxt][y].priority > d {
				pq.update(items[nxt][y], d)
			}
		}

		if y > 0 {
			dd := getDist(y-1, nxt-1) + d
			if items[x][nxt].priority > dd {
				pq.update(items[x][nxt], dd)
			}
		} else {
			if items[x][nxt].priority > d {
				pq.update(items[x][nxt], d)
			}
		}
	}

	return 0
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

// An Item is something we manage in a priority queue.
type Item struct {
	value    int // The value of the item; arbitrary.
	priority int // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) update(item *Item, priority int) {
	// if item.index < 0 {
	// 	return
	// }
	item.priority = priority
	heap.Fix(pq, item.index)
}
