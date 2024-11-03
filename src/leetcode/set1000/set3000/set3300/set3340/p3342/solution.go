package p3342

import "container/heap"

const inf = 1 << 60

var dd = []int{-1, 0, 1, 0, -1}

func minTimeToReach(moveTime [][]int) int {
	n := len(moveTime)
	m := len(moveTime[0])
	dp := make([][][]*Item, n)
	pq := make(PriorityQueue, n*m*2)
	for i := 0; i < n; i++ {
		dp[i] = make([][]*Item, m)
		for j := 0; j < m; j++ {
			dp[i][j] = make([]*Item, 2)
			for k := 0; k < 2; k++ {
				it := new(Item)
				it.value = i*2*m + j*2 + k
				it.priority = inf
				it.index = it.value
				dp[i][j][k] = it
				pq[it.index] = it
			}
		}
	}

	heap.Init(&pq)
	pq.update(dp[0][0][0], 0)

	for pq.Len() > 0 {
		it := heap.Pop(&pq).(*Item)
		k := it.value % 2
		j := (it.value / 2) % m
		i := it.value / 2 / m

		for x := 0; x < 4; x++ {
			r, c := i+dd[x], j+dd[x+1]
			if r < 0 || r == n || c < 0 || c == m {
				continue
			}
			if dp[r][c][k^1].priority > max(it.priority, moveTime[r][c])+k+1 {
				pq.update(dp[r][c][k^1], max(it.priority, moveTime[r][c])+k+1)
			}
		}
	}
	res := min(dp[n-1][m-1][0].priority, dp[n-1][m-1][1].priority)
	return res
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

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
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
	// item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}
