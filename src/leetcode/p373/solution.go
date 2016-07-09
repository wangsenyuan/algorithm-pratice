package main

import (
	"container/heap"
	"fmt"
)

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	if len(nums1) == 0 || len(nums2) == 0 {
		return nil
	}
	result := make([][]int, 0, k)

	visited := make([][]bool, len(nums1))
	for i := range visited {
		visited[i] = make([]bool, len(nums2))
	}

	pq := make(PriorityQueue, 1)
	pq[0] = &item{0, 0, nums1[0] + nums2[0]}

	heap.Init(&pq)
	for i := 0; i < k && pq.Len() > 0; i++ {
		x := heap.Pop(&pq).(*item)
		result = append(result, []int{nums1[x.i], nums2[x.j]})
		add(&pq, nums1, nums2, x.i, x.j+1, visited)
		add(&pq, nums1, nums2, x.i+1, x.j, visited)
	}

	return result
}

func add(pq *PriorityQueue, nums1 []int, nums2 []int, i int, j int, visited [][]bool) {
	if i >= len(nums1) || j >= len(nums2) || visited[i][j] {
		return
	}

	heap.Push(pq, &item{i, j, nums1[i] + nums2[j]})
	visited[i][j] = true
}

type item struct {
	i, j, x int
}

type PriorityQueue []*item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].x < pq[j].x
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*item)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func main() {
	nums1 := []int{1, 1, 2}
	nums2 := []int{1, 2, 3}
	result := kSmallestPairs(nums1, nums2, 8)
	for _, x := range result {
		fmt.Printf("%v\n", x)
	}
}
