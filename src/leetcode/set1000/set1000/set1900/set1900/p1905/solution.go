package p1905

import "sort"

func minDifference(nums []int, queries [][]int) []int {
	tree := NewTree(101)

	qs := make([]Q, len(queries))

	for i := 0; i < len(queries); i++ {
		qs[i] = Q{queries[i][0], queries[i][1], i}
	}

	sort.Slice(qs, func(i, j int) bool {
		a, b := qs[i].r/BLOCK, qs[j].r/BLOCK
		if a < b {
			return true
		}
		if a > b {
			return false
		}
		return qs[i].l < qs[j].l
	})

	ans := make([]int, len(qs))

	L, R := 0, 0

	for _, q := range qs {
		l, r := q.l, q.r
		r++
		for r > R {
			tree.Update(nums[R], 1)
			R++
		}
		for r < R {
			R--
			tree.Update(nums[R], -1)
		}

		for l < L {
			L--
			tree.Update(nums[L], 1)
		}

		for l > L {
			tree.Update(nums[L], -1)
			L++
		}
		ans[q.idx] = tree.val[1]
		if ans[q.idx] >= INF {
			ans[q.idx] = -1
		}
	}

	return ans
}

type Q struct {
	l, r int
	idx  int
}

const BLOCK = 200

type Tree struct {
	min []int
	max []int
	on  []int
	val []int
	m   int
}

func NewTree(n int) *Tree {
	m := 1
	for m < n {
		m <<= 1
	}
	min := make([]int, m<<1)
	max := make([]int, m<<1)
	on := make([]int, m<<1)
	val := make([]int, m<<1)
	for i := 0; i < len(val); i++ {
		val[i] = INF
	}
	return &Tree{min, max, on, val, m}
}

const INF = 1 << 30

func (t *Tree) Update(p int, v int) {
	merge := func(i int) {
		t.on[i] = t.on[2*i] + t.on[2*i+1]
		if t.on[2*i] == 0 {
			t.min[i] = t.min[2*i+1]
			t.max[i] = t.max[2*i+1]
			t.val[i] = t.val[2*i+1]
			return
		}
		if t.on[2*i+1] == 0 {
			t.min[i] = t.min[2*i]
			t.max[i] = t.max[2*i]
			t.val[i] = t.val[2*i]
			return
		}
		t.min[i] = t.min[2*i]
		t.max[i] = t.max[2*i+1]
		t.val[i] = t.val[2*i]
		if t.val[i] > t.val[2*i+1] {
			t.val[i] = t.val[2*i+1]
		}
		if t.val[i] > t.min[2*i+1]-t.max[2*i] {
			t.val[i] = t.min[2*i+1] - t.max[2*i]
		}
	}

	var loop func(i int, l, r int)
	loop = func(i int, l, r int) {
		if l == r {
			t.min[i] = p
			t.max[i] = p
			t.on[i] += v
			//t.val[l] = INF
			return
		}
		mid := (l + r) / 2
		if p <= mid {
			loop(i*2, l, mid)
		} else {
			loop(i*2+1, mid+1, r)
		}
		merge(i)
	}
	loop(1, 0, t.m-1)
}
