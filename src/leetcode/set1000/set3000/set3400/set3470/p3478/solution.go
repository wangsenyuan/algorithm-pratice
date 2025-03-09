package p3478

import (
	"container/heap"
	"slices"
)

func findMaxSum(nums1 []int, nums2 []int, k int) []int64 {
	n := len(nums1)
	type pair struct {
		first  int
		second int
	}
	arr := make([]pair, n)
	for i, num := range nums1 {
		arr[i] = pair{num, i}
	}
	slices.SortFunc(arr, func(a, b pair) int {
		return a.first - b.first
	})

	ans := make([]int64, n)
	pq := make(IntHeap, 0, n)
	var sum int
	for i := 0; i < n; {
		j := i
		for i < n && arr[i].first == arr[j].first {
			ans[arr[i].second] = int64(sum)
			i++
		}
		for j < i {
			sum += nums2[arr[j].second]
			heap.Push(&pq, nums2[arr[j].second])
			if pq.Len() > k {
				sum -= heap.Pop(&pq).(int)
			}
			j++
		}
	}

	return ans
}

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
