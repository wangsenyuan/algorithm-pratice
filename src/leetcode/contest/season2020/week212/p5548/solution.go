package p5548

import "container/heap"

const INF = 1000000000

var dd = []int{-1, 0, 1, 0, -1}

func minimumEffortPath(heights [][]int) int {
	m := len(heights)
	n := len(heights[0])
	cells := make([][]*Item, m)
	pq := make(PriorityQueue, 0, m*n)
	for i := 0; i < m; i++ {
		cells[i] = make([]*Item, n)
		for j := 0; j < n; j++ {
			cell := new(Item)
			cell.x = i
			cell.y = j
			cell.priority = INF
			heap.Push(&pq, cell)
			cells[i][j] = cell
		}
	}

	pq.update(cells[0][0], 0)

	for pq.Len() > 0 {
		cur := heap.Pop(&pq).(*Item)
		x, y := cur.x, cur.y
		if x == m-1 && y == n-1 {
			break
		}
		for k := 0; k < 4; k++ {
			u, v := x+dd[k], y+dd[k+1]
			if u >= 0 && u < m && v >= 0 && v < n && cells[u][v].index >= 0 && cells[u][v].priority > max(cells[x][y].priority, abs(heights[u][v]-heights[x][y])) {
				pq.update(cells[u][v], max(cells[x][y].priority, abs(heights[u][v]-heights[x][y])))
			}
		}
	}

	return cells[m-1][n-1].priority
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
	x, y     int
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
	item.priority = priority
	heap.Fix(pq, item.index)
}
