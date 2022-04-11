package p2232

import "container/heap"

const MOD = 1000000007

func maximumProduct(nums []int, k int) int {
	n := len(nums)

	if n == 1 {
		return nums[0] + k
	}

	pq := IntHeap(nums)

	heap.Init(&pq)

	for k > 0 {
		x := heap.Pop(&pq).(int)
		y := pq[0]
		if x == y {
			x++
			k--
			heap.Push(&pq, x)
			continue
		}
		d := y - x
		if d <= k {
			x += d
			k -= d
			heap.Push(&pq, x)
		} else {
			// y - x > k
			x += k
			k = 0
			heap.Push(&pq, x)
		}
	}

	var res int64 = 1

	for i := 0; i < len(pq); i++ {
		res *= int64(pq[i])
		res %= MOD
	}

	return int(res)
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
