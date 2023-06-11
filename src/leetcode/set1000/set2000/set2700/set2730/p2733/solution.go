package p2733

import "sort"

func maximumSumQueries(nums1 []int, nums2 []int, queries [][]int) []int {
	ys := make(map[int]int)

	for _, num := range nums2 {
		ys[num]++
	}

	for _, qs := range queries {
		ys[qs[1]]++
	}

	arr := make([]int, 0, len(ys))
	for y := range ys {
		arr = append(arr, y)
	}

	sort.Ints(arr)

	for i := 0; i < len(arr); i++ {
		ys[arr[i]] = i
	}

	nums := make([]Pair, len(nums1))
	for i := 0; i < len(nums1); i++ {
		nums[i] = Pair{nums1[i], nums2[i]}
	}

	qs := make([]Query, len(queries))

	for i := 0; i < len(queries); i++ {
		qs[i] = Query{queries[i][0], queries[i][1], i}
	}

	sort.Slice(nums, func(i, j int) bool {
		return nums[i].first > nums[j].first
	})

	sort.Slice(qs, func(i, j int) bool {
		return qs[i].x > qs[j].x
	})
	ans := make([]int, len(qs))

	tr := NewTree(len(ys))

	for i, j := 0, 0; i < len(qs); i++ {
		for j < len(nums) && nums[j].first >= qs[i].x {
			x, y := nums[j].first, nums[j].second
			id := ys[y]
			tr.Update(id, x+y)
			j++
		}
		ans[qs[i].id] = tr.Get(ys[qs[i].y], len(ys))
	}

	return ans
}

type Tree struct {
	arr []int
	n   int
}

func NewTree(n int) *Tree {
	arr := make([]int, 2*n)
	for i := 0; i < len(arr); i++ {
		arr[i] = -1
	}
	return &Tree{arr, n}
}

func (t *Tree) Update(p int, v int) {
	p += t.n
	t.arr[p] = max(t.arr[p], v)
	for p > 1 {
		t.arr[p>>1] = max(t.arr[p], t.arr[p^1])
		p >>= 1
	}
}

func (t *Tree) Get(l int, r int) int {
	l += t.n
	r += t.n
	var res int = -1
	for l < r {
		if l&1 == 1 {
			res = max(res, t.arr[l])
			l++
		}
		if r&1 == 1 {
			r--
			res = max(res, t.arr[r])
		}
		l >>= 1
		r >>= 1
	}
	return res
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

type Query struct {
	x  int
	y  int
	id int
}

type Pair struct {
	first  int
	second int
}
