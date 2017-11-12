package main

import (
	"fmt"
	"sort"
)

type interval struct {
	start  int
	end    int
	height int
}

func fallingSquares1(positions [][]int) []int {
	n := len(positions)
	res := make([]int, n)
	intervals := make([]interval, 0, n)
	best := 0

	for i := 0; i < n; i++ {
		cur := positions[i]
		left, side := cur[0], cur[1]
		height := side

		for _, inter := range intervals {
			if inter.end <= left || left+side <= inter.start {
				continue
			}
			if inter.height+side > height {
				height = inter.height + side
			}
		}

		intervals = append(intervals, interval{left, left + side, height})

		if height > best {
			best = height
		}
		res[i] = best
	}

	return res
}

type SegmentTree struct {
	n    int
	h    int
	tree []int
	lazy []int
}

func InitSegmentTree(n int) *SegmentTree {
	h := 1
	for 1<<uint(h) < n {
		h++
	}
	return &SegmentTree{n, h, make([]int, 2*n), make([]int, n)}
}

func (st *SegmentTree) apply(i, v int) {
	st.tree[i] = max(st.tree[i], v)
	if i < st.n {
		st.lazy[i] = max(st.lazy[i], v)
	}
}

func (st *SegmentTree) pull(i int) {
	for i > 1 {
		i >>= 1
		st.tree[i] = max(st.tree[2*i], st.tree[2*i+1])
		st.tree[i] = max(st.tree[i], st.lazy[i])
	}
}

func (st *SegmentTree) Update(left int, right int, h int) {
	left += st.n
	right += st.n
	i, j := left, right
	for left <= right {
		if left&1 == 1 {
			st.apply(left, h)
			left++
		}
		if right&1 == 0 {
			st.apply(right, h)
			right--
		}
		left >>= 1
		right >>= 1
	}
	st.pull(i)
	st.pull(j)
}

func (st *SegmentTree) Query(left int, right int) int {
	left += st.n
	right += st.n
	ans := 0
	st.push(left)
	st.push(right)
	for left <= right {
		if left&1 == 1 {
			ans = max(ans, st.tree[left])
			left++
		}
		if right&1 == 0 {
			ans = max(ans, st.tree[right])
			right--
		}
		left >>= 1
		right >>= 1
	}
	return ans
}

func (st *SegmentTree) push(i int) {
	for j := st.h; j > 0; j-- {
		k := i >> uint(j)
		if st.lazy[k] > 0 {
			st.apply(k*2, st.lazy[k])
			st.apply(k*2+1, st.lazy[k])
			st.lazy[k] = 0
		}
	}

}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func fallingSquares(positions [][]int) []int {
	set := make(map[int]bool)

	for _, pos := range positions {
		set[pos[0]] = true
		set[pos[0]+pos[1]-1] = true
	}
	n := len(set)

	arr := make([]int, n)

	idx := 0
	for coord, _ := range set {
		arr[idx] = coord
		idx++
	}

	sort.Ints(arr)

	index := make(map[int]int)

	for i, coord := range arr {
		index[coord] = i
	}

	st := InitSegmentTree(n)
	ans := make([]int, len(positions))
	best := 0
	for i, pos := range positions {
		left := index[pos[0]]
		right := index[pos[0]+pos[1]-1]
		height := st.Query(left, right) + pos[1]
		st.Update(left, right, height)
		best = max(best, height)
		ans[i] = best
	}

	return ans
}

func main() {
	fmt.Println(fallingSquares([][]int{{1, 2}, {2, 3}, {6, 1}}))
	fmt.Println(fallingSquares([][]int{{100, 100}, {200, 100}}))
}
