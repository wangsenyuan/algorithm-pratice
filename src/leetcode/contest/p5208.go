package contest

import (
	"container/heap"
	"math"
)

func minimumMoves(grid [][]int) int {
	n := len(grid)
	N := n * n * 2
	items := make([]*Item, N)

	pq := make(PriorityQueue, 0, N)
	heap.Init(&pq)

	for i := 0; i < N; i++ {
		item := new(Item)
		item.value = i
		item.priority = math.MaxInt32
		items[i] = item
		heap.Push(&pq, item)
	}

	pq.update(items[toPos(0, 1, 0, n)], 0)

	for pq.Len() > 0 {
		curItem := heap.Pop(&pq).(*Item)
		cur := curItem.value

		if curItem.priority == math.MaxInt32 {
			break
		}

		x, y, d := fromPos(cur, n)

		if d == 0 {
			// horizontal
			// try right
			if y < n-1 && grid[x][y+1] == 0 {
				// empty
				tmp := toPos(x, y+1, 0, n)
				if items[tmp].priority > items[cur].priority+1 {
					pq.update(items[tmp], items[cur].priority+1)
				}
			}
			if x < n-1 && grid[x+1][y-1] == 0 && grid[x+1][y] == 0 {
				// try down
				tmp := toPos(x+1, y, 0, n)
				if tmp >= 0 {
					if items[tmp].priority > items[cur].priority+1 {
						pq.update(items[tmp], items[cur].priority+1)
					}
				}
				// try rotate clockwise
				tmp = toPos(x+1, y-1, 1, n)
				if tmp >= 0 {
					if items[tmp].priority > items[cur].priority+1 {
						pq.update(items[tmp], items[cur].priority+1)
					}
				}
			}
		} else {
			if y < n-1 && grid[x-1][y+1] == 0 && grid[x][y+1] == 0 {
				// try right
				tmp := toPos(x, y+1, 1, n)
				if tmp >= 0 {
					if items[tmp].priority > items[cur].priority+1 {
						pq.update(items[tmp], items[cur].priority+1)
					}
				}
				// try rotate
				tmp = toPos(x-1, y+1, 0, n)
				if tmp >= 0 {
					if items[tmp].priority > items[cur].priority+1 {
						pq.update(items[tmp], items[cur].priority+1)
					}
				}
			}
			if x < n-1 && grid[x+1][y] == 0 {
				// try down
				tmp := toPos(x+1, y, 1, n)
				if tmp >= 0 {
					if items[tmp].priority > items[cur].priority+1 {
						pq.update(items[tmp], items[cur].priority+1)
					}
				}
			}
		}
	}

	res := items[toPos(n-1, n-1, 0, n)]
	if res.priority == math.MaxInt32 {
		return -1
	}
	return res.priority
}

func fromPos(pos int, n int) (int, int, int) {
	d := pos % 2
	y := (pos / 2) % n
	x := (pos / 2) / n
	return x, y, d
}

func toPos(x, y, d int, n int) int {
	if x < 0 || x >= n || y < 0 || y >= n || d < 0 || d >= 2 {
		return -1
	}
	return (x*n+y)*2 + d
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
	if item.index < 0 {
		return
	}
	item.priority = priority
	heap.Fix(pq, item.index)
}
