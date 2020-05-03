package p1439

import (
	"container/heap"
	"sort"
)

const INF = 5001

func kthSmallest(mat [][]int, k int) int {

	getFirstKPairs := func(a, b []int, k int) []int {
		// a & b are sorted
		if len(a) == 0 || len(b) == 0 || k == 0 {
			return nil
		}
		pq := make(PriorityQueue, 0, len(a)*len(b))

		createItem := func(i, j int) *Item {
			item := new(Item)
			item.pos1 = i
			item.pos2 = j
			// item.pos = j
			item.priority = a[i] + b[j]
			return item
		}

		for j := 0; j < len(b); j++ {
			heap.Push(&pq, createItem(0, j))
		}
		res := make([]int, 0, min(k, len(a)*len(b)))

		for pq.Len() > 0 && k > 0 {
			cur := heap.Pop(&pq).(*Item)
			k--
			res = append(res, cur.priority)

			if cur.pos1+1 < len(a) {
				heap.Push(&pq, createItem(cur.pos1+1, cur.pos2))
			}
		}

		return res
	}

	var loop func(arr [][]int) []int

	loop = func(arr [][]int) []int {
		if len(arr) == 0 {
			return nil
		}
		if len(arr) == 1 {
			return arr[0]
		}
		mid := len(arr) / 2
		a := loop(arr[:mid])
		b := loop(arr[mid:])
		return getFirstKPairs(a, b, k)
	}

	res := loop(mat)

	return res[k-1]
}

// An Item is something we manage in a priority queue.
type Item struct {
	pos1, pos2 int // The value of the item; arbitrary.
	priority   int // The priority of the item in the queue.
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

func kthSmallest2(mat [][]int, k int) int {

	var loop func(arr [][]int) []int

	loop = func(arr [][]int) []int {
		if len(arr) == 0 {
			return []int{0}
		}
		if len(arr) == 1 {
			return arr[0]
		}

		mid := len(arr) / 2

		a := loop(arr[:mid])
		b := loop(arr[mid:])

		c := make([]int, len(a)*len(b))

		for i := 0; i < len(a); i++ {
			for j := 0; j < len(b); j++ {
				c[i*len(b)+j] = a[i] + b[j]
			}
		}

		sort.Ints(c)

		if len(c) > k {
			return c[:k]
		}
		return c
	}

	res := loop(mat)

	return res[k-1]
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
func kthSmallest1(mat [][]int, k int) int {
	n := len(mat)
	// INF := k + 1

	mem := make([]map[int]int, n)

	for i := 0; i < n; i++ {
		mem[i] = make(map[int]int)
	}

	var count func(sum int, row int) int

	// count how many numbers that is less than or equal to sum

	count = func(sum int, row int) int {
		if sum < 0 {
			return 0
		}
		if row == n {
			return 1
		}

		if v, found := mem[row][sum]; found {
			return v
		}

		var res int

		for i := 0; i < len(mat[row]) && mat[row][i] <= sum; i++ {
			res += count(sum-mat[row][i], row+1)
			if res > k {
				break
			}
		}

		mem[row][sum] = res

		return res
	}

	var right int

	for i := 0; i < n; i++ {
		right += mat[i][len(mat[i])-1]
	}

	var left int
	right++

	for left < right {
		mid := (left + right) / 2
		c := count(mid, 0)
		if c >= k {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return right
}
