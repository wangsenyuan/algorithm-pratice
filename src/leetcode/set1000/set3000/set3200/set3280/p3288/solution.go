package p3288

import (
	"slices"
	"sort"
)

func maxPathLength(coordinates [][]int, k int) int {
	var bl [][]int
	var tr [][]int
	a := coordinates[k]
	for _, coord := range coordinates {
		if coord[0] < a[0] && coord[1] < a[1] {
			bl = append(bl, coord)
		}
		if coord[0] > a[0] && coord[1] > a[1] {
			tr = append(tr, coord)
		}
	}

	return 1 + solve(bl) + solve(tr)
}

func solve(arr [][]int) int {
	if len(arr) == 0 {
		return 0
	}
	n := len(arr)
	vs := make([]int, n)
	for i, v := range arr {
		vs[i] = v[1]
	}
	sort.Ints(vs)
	var ys []int
	for i := 0; i < n; {
		ys = append(ys, vs[i])
		y := vs[i]
		for i < n && vs[i] == y {
			i++
		}
	}

	set := NewSegTree(len(ys))

	slices.SortFunc(arr, func(a, b []int) int {
		if a[0] != b[0] {
			return a[0] - b[0]
		}
		// 在相同x时，先处理高处的，避免被影响
		return b[1] - a[1]
	})

	dp := make([]int, len(ys))

	for _, cur := range arr {
		y := cur[1]
		y = sort.SearchInts(ys, y)
		v := set.Query(0, y)
		if v+1 > dp[y] {
			dp[y] = v + 1
			set.Update(y, v+1)
		}
	}
	return slices.Max(dp)
}

type SegTree struct {
	arr  []int
	size int
}

func NewSegTree(n int) *SegTree {
	arr := make([]int, 2*n)
	size := n
	return &SegTree{arr, size}
}

func (st *SegTree) Update(pos int, val int) {
	arr := st.arr
	n := st.size
	arr[pos+n] = val
	pos += n
	for i := pos; i > 1; i >>= 1 {
		arr[i>>1] = max(arr[i], arr[i^1])
	}
}

func (st *SegTree) Query(left, right int) int {
	arr := st.arr
	n := st.size
	left += n
	right += n
	var res int
	for left < right {
		if left&1 == 1 {
			res = max(res, arr[left])
			left++
		}
		if right&1 == 1 {
			right--
			res = max(res, arr[right])
		}
		left >>= 1
		right >>= 1
	}
	return res
}
