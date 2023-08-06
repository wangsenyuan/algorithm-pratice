package p2810

import (
	"container/heap"
	"sort"
)

func findMaximumElegance(items [][]int, k int) int64 {
	sort.Slice(items, func(i, j int) bool {
		return items[i][0] > items[j][0]
	})

	n := len(items)

	var sum int64
	vis := make(map[int]int)

	pq := make(IntHeap, 0, n)

	for i := 0; i < k; i++ {
		sum += int64(items[i][0])
		vis[items[i][1]]++
		if vis[items[i][1]] > 1 {
			heap.Push(&pq, items[i][0])
		}
	}

	ans := sum + int64(len(vis))*int64(len(vis))

	for i := k; i < n && pq.Len() > 0; i++ {
		cur := items[i]
		if vis[cur[1]] > 0 {
			// category already in the result, not good to use this one
			continue
		}
		//  ans - x * x - a + (x + 1) * (x + 1) + b
		sum += int64(cur[0] - pq[0])
		heap.Pop(&pq)
		vis[cur[1]]++
		ans = max(ans, sum+int64(len(vis))*int64(len(vis)))
	}

	return ans
}

func max(a, b int64) int64 {
	if a >= b {
		return a
	}
	return b
}

// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
