package p1494

import (
	"container/heap"
	"math"
)

func minNumberOfSemesters(n int, dependencies [][]int, k int) int {
	N := 1 << n
	pre := make([]int, n)

	for _, dep := range dependencies {
		u, v := dep[0], dep[1]
		u--
		v--
		pre[v] |= 1 << u
	}

	set := make([]int, N)
	valid := make([]bool, N)

	for mask := 1; mask < N; mask++ {
		if digitCount(mask) <= k {
			for i := 0; i < n; i++ {
				if mask&(1<<i) > 0 {
					set[mask] |= pre[i]
				}
			}
			valid[mask] = set[mask]&mask == 0
		}
	}

	dp := make([]int, N)
	dp[0] = 0
	for i := 1; i < N; i++ {
		dp[i] = math.MaxInt32
	}
	for mask := 1; mask < N; mask++ {
		for sub := mask; sub > 0; sub = (sub - 1) & mask {
			if valid[sub] && mask&set[sub] == set[sub] {
				dp[mask] = min(dp[mask], dp[mask^sub]+1)
			}
		}
	}

	return dp[N-1]
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func digitCount(num int) int {
	var res int

	for num > 0 {
		res += num & 1
		num >>= 1
	}
	return res
}

func minNumberOfSemesters1(n int, dependencies [][]int, k int) int {
	degree := make([]int, n)
	out := make([][]int, n)
	for _, dep := range dependencies {
		degree[dep[1]-1]++
		u := dep[0] - 1
		if out[u] == nil {
			out[u] = make([]int, 0, 10)
		}
		out[u] = append(out[u], dep[1]-1)
	}

	items := make([]*Item, n)
	pq := make(PriorityQueue, n)

	for i := 0; i < n; i++ {
		item := new(Item)
		item.pos = i
		item.priority = degree[i]
		item.priority2 = len(out[i])
		item.index = i
		items[i] = item
		pq[i] = item
	}

	heap.Init(&pq)

	var res int

	for pq.Len() > 0 {
		ii := make([]*Item, 0, k)

		var cnt int
		for pq.Len() > 0 && pq[0].priority == 0 && cnt < k {
			cnt++
			ii = append(ii, heap.Pop(&pq).(*Item))
		}
		res++

		for _, cur := range ii {
			u := cur.pos

			for _, v := range out[u] {
				tmp := items[v]
				// tmp.priority--
				pq.update(tmp, tmp.priority-1)
			}
		}
	}

	return res
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

// An Item is something we manage in a priority queue.
type Item struct {
	pos       int // The value of the item; arbitrary.
	priority  int // The priority of the item in the queue.
	priority2 int
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].priority < pq[j].priority || pq[i].priority == pq[j].priority && pq[i].priority2 > pq[j].priority2
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
