package p1540

import "container/heap"

func canConvertString(s string, t string, k int) bool {
	if len(s) != len(t) {
		return false
	}
	mem := make(map[int]int)
	var cnt int
	for i := 0; i < len(s); i++ {
		x := int(t[i]-'a') - int(s[i]-'a')
		if x == 0 {
			continue
		}
		if x < 0 {
			x += 26
		}
		mem[x]++
		cnt = max(cnt, x+26*(mem[x]-1))
	}
	return cnt <= k
}
func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func canConvertString1(s string, t string, k int) bool {
	if len(s) != len(t) {
		return false
	}
	n := len(s)

	pq := make(IntHeap, 0, n)

	for i := 0; i < n; i++ {
		x := int(t[i]-'a') - int(s[i]-'a')
		if x == 0 {
			continue
		}
		if x < 0 {
			x += 26
		}
		if x > k {
			return false
		}
		heap.Push(&pq, x)
	}

	j := 1
	for pq.Len() > 0 && j <= k {
		cur := heap.Pop(&pq).(int)
		if j <= cur {
			j = cur + 1
			continue
		}
		for cur+26 <= j {
			cur += 26
		}

		heap.Push(&pq, cur+26)
	}
	return pq.Len() == 0 && j <= k+1
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
