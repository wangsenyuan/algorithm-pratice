package main

import (
	"container/heap"
	"math"
)

func main() {

}

func solve(R, C, F int, G []string) int {
	holes := make([][]int, 0, R*C)
	var index int
	// the hole index for (r, c)
	X := make([][]int, R)
	for i := 0; i < R; i++ {
		X[i] = make([]int, C)
		for j := 0; j < C; j++ {
			X[i][j] = -1
		}
	}
	bricks := make([][]int, R*C)

	for i := 0; i < R; i++ {
		bricks[i] = make([]int, C)
		for j := 0; j < C; j++ {
			if G[i][j] == '.' {
				//a hole
				if j == 0 || G[i][j-1] == '#' {
					// the start hole
					holes[index] = make([]int, 3)
					holes[index][0] = i
					holes[index][1] = j
				}
				X[i][j] = index
			} else {
				bricks[i][j]++
				if i > 0 {
					bricks[i][j] += bricks[i][j-1]
				}
				if j > 0 && G[i][j-1] == '.' {
					// the end hole
					holes[index][2] = j
					index++
				}
			}
		}
	}

	//bfs to reach every hole
	que := make(PriorityQueue, 0, index)
	items := make([]*Item, index)
	for i := 0; i < index; i++ {
		item := new(Item)
		item.pos = i
		item.dist = math.MaxInt32
		items[i] = item
		heap.Push(&que, item)
	}

	que.update(items[0], 0)

	for que.Len() > 0 {
		cur := heap.Pop(&que).(*Item)
		i := cur.pos
		row, left, right := holes[i][0], holes[i][1], holes[i][2]
		if row == R-1 && right == C-1 {
			//the destination reach
			return cur.dist
		}

		dig := func(row1, col1, row2, col2 int) int {
			a := abs(row2 - row1)
			b := abs(col2 - col1)
			if b < a {
				return -1
			}
			var res int
			if col1 < col2 {
				//dig afterward
				c := col2 - b
				for r := row1 + 1; r < row2; r++ {
					c++
					if G[r][c] == '.' {
						//a dangerous hole
						return -1
					}
					res += bricks[r][c] - bricks[row1][c]
				}
			} else {
				//dig forward
				c := col2 + b
				for r := row1 + 1; r < row2; r++ {
					c--
					if G[r][c] == '.' {
						return -1
					}
					res += bricks[r][c] - bricks[row1][c]
				}
			}

			return res
		}

		for i := 0; i < index; i++ {
			if items[i].index < 0 {
				// already out of que
				continue
			}
			next := items[i]
			j := next.pos
			nrow, nleft, nright := holes[j][0], holes[j][1], holes[j][2]
			if nrow <= row || right < nleft || left > nright {
				// can't reach this one
				continue
			}
			a := max(left, nleft)
			b := min(right, nright)
			best := -1
			for k := a; k <= b; k++ {
				//to reach position (nrow - 1, k)
				cnt := dig(row, left, nrow-1, k)
				if cnt >= 0 {
					// from (nrow-1, k)
					fall := 1
					for u := nrow; u < R && G[u][k] == '.'; u++ {
						fall++
					}
					if fall < F && (best < 0 || best > cnt) {
						best = cnt
					}
				}
			}
			if best >= 0 {
				que.update(items[j], best)
			}
		}
	}
	return -1
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

// An Item is something we manage in a priority queue.
type Item struct {
	pos  int // The value of the item; arbitrary.
	dist int // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].dist < pq[j].dist
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
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) update(item *Item, dist int) {
	item.dist = dist
	heap.Fix(pq, item.index)
}
