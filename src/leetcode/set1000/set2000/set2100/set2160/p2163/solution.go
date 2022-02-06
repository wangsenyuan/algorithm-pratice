package p2163

import (
	"container/heap"
	"math"
)

func minimumDifference(nums []int) int64 {
	N := len(nums)
	n := N / 3
	// 如果一个数比较大，如果它属于前半部分，则需要删除，如果属于后半部分，则需要保留
	sum := make([]int64, N)

	right := make(MinHeap, 0, n)
	var tot int64
	for i := N - 1; i >= n; i-- {
		if right.Len() < n {
			tot += int64(nums[i])
			heap.Push(&right, nums[i])
		} else if nums[i] > right[0] {
			tot -= int64(heap.Pop(&right).(int))
			tot += int64(nums[i])
			heap.Push(&right, nums[i])
		}
		sum[i] = tot
	}

	left := make(MaxHeap, 0, n)
	tot = 0
	var best int64 = math.MaxInt64
	for i := 0; i < 2*n; i++ {
		if left.Len() < n {
			tot += int64(nums[i])
			heap.Push(&left, nums[i])
		} else if nums[i] < left[0] {
			tot -= int64(heap.Pop(&left).(int))
			tot += int64(nums[i])
			heap.Push(&left, nums[i])
		}
		
		if left.Len() == n {
			best = min(best, tot-sum[i+1])
		}
	}
	return best
}

func min(a, b int64) int64 {
	if a <= b {
		return a
	}
	return b
}

// An IntHeap is a min-heap of ints.
type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// An IntHeap is a min-heap of ints.
type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
