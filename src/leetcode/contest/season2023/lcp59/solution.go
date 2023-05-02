package lcp59

import "container/heap"

func buildBridge(num int, wood [][]int) int64 {
	// 头疼
	L := NewHeap(greater)
	R := NewHeap(less)
	heap.Push(L, int64(wood[0][0]))
	heap.Push(R, int64(wood[0][0]))

	var l, r, ans int64

	for i := 1; i < len(wood); i++ {
		l -= int64(wood[i][1] - wood[i][0])
		r += int64(wood[i-1][1] - wood[i-1][0])
		a := L.arr[0] + l
		b := R.arr[0] + r
		x := int64(wood[i][0])
		if x < a {
			ans += a - x
			heap.Pop(L)
			heap.Push(L, x-l)
			heap.Push(L, x-l)
			heap.Push(R, a-r)
		} else if x > b {
			ans += x - b
			heap.Pop(R)
			heap.Push(R, x-r)
			heap.Push(R, x-r)
			heap.Push(L, b-l)
		} else {
			heap.Push(L, x-l)
			heap.Push(R, x-r)
		}
	}

	return ans
}

const INF = 1 << 60

func min(a, b int64) int64 {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int64) int64 {
	if a >= b {
		return a
	}
	return b
}

func less(a int64, b int64) bool {
	return min(a, b) == a
}

func greater(a int64, b int64) bool {
	return max(a, b) == a
}

type LongHeap struct {
	arr []int64
	cmp func(int64, int64) bool
}

func NewHeap(f func(int64, int64) bool) *LongHeap {
	pq := new(LongHeap)
	pq.cmp = f
	pq.arr = make([]int64, 0, 1)
	return pq
}

func (this *LongHeap) Len() int {
	return len(this.arr)
}

func (this *LongHeap) Less(i, j int) bool {
	return this.cmp(this.arr[i], this.arr[j])
}

func (this *LongHeap) Swap(i, j int) {
	this.arr[i], this.arr[j] = this.arr[j], this.arr[i]
}

func (this *LongHeap) Push(x interface{}) {
	i := x.(int64)
	this.arr = append(this.arr, i)
}

func (this *LongHeap) Pop() interface{} {
	old := this.arr
	n := len(old)
	res := old[n-1]
	this.arr = old[:n-1]
	return res
}
