package p2928

import "container/heap"

func maxSpending(values [][]int) int64 {
	// 一个观察是，商品价格越高，越早买越好，
	// 那是不是每次都购买最便宜的那个，就是最优的呢？
	// a, b, c, ..., x, y, z
	// a <= b <= c <= ... <= x <= y <= z
	m := len(values)
	n := len(values[0])

	pq := make(PQ, 0, m)
	for i := 0; i < m; i++ {
		item := Item{i, n - 1, values[i][n-1]}
		heap.Push(&pq, item)
	}

	var res int64

	for d := 1; d <= m*n; d++ {
		cur := heap.Pop(&pq).(Item)

		res += int64(d) * int64(cur.value)
		i, j := cur.i, cur.j
		if j > 0 {
			it := Item{i, j - 1, values[i][j-1]}
			heap.Push(&pq, it)
		}
	}
	return res
}

type Item struct {
	i     int
	j     int
	value int
}

// An IntHeap is a min-heap of ints.
type PQ []Item

func (h PQ) Len() int           { return len(h) }
func (h PQ) Less(i, j int) bool { return h[i].value < h[j].value }
func (h PQ) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *PQ) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(Item))
}

func (h *PQ) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
