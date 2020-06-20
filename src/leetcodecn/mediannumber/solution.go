package mediannumber

import "container/heap"

type MedianFinder struct {
	min *MinHeap
	max *MaxHeap
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	mh := make(MinHeap, 0, 100)
	xh := make(MaxHeap, 0, 100)
	return MedianFinder{&mh, &xh}
}

func (this *MedianFinder) AddNum(num int) {
	if this.min.Len() > 0 && (*this.min)[0] < num {
		heap.Push(this.min, num)
	} else {
		heap.Push(this.max, num)
	}

	for this.min.Len() > this.max.Len() {
		heap.Push(this.max, heap.Pop(this.min).(int))
	}

	for this.min.Len() < this.max.Len() {
		heap.Push(this.min, heap.Pop(this.max).(int))
	}
}

func (this *MedianFinder) FindMedian() float64 {
	n := this.min.Len() + this.max.Len()
	var res float64
	if n%2 == 0 {
		res = float64((*this.min)[0]+(*this.max)[0]) / 2
	} else {
		res = float64((*this.min)[0])
	}

	return res
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

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
