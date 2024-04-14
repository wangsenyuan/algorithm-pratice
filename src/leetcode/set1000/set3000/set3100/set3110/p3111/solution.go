package p3111

import "sort"

func numberOfSubarrays(nums []int) int64 {
	n := len(nums)
	id := make([]int, n)
	for i := 0; i < n; i++ {
		id[i] = i
	}
	sort.Slice(id, func(i, j int) bool {
		return nums[id[i]] > nums[id[j]] || nums[id[i]] == nums[id[j]] && id[i] < id[j]
	})

	tr := NewSegTree(n, 0, plus)
	var res int
	for i := 0; i < n; {
		j := i
		var cnt int
		for i < n && nums[id[i]] == nums[id[j]] {
			if i > j {
				sum := tr.Get(id[i-1], id[i])
				if sum > 0 {
					cnt = 0
				}
			}
			cnt++

			res += cnt
			i++
		}
		for j < i {
			tr.Update(id[j], 1)
			j++
		}
	}
	return int64(res)
}

func plus(a, b int) int {
	return a + b
}

type SegTree struct {
	size       int
	arr        []int
	init_value int
	op         func(int, int) int
}

func NewSegTree(n int, v int, op func(int, int) int) *SegTree {
	arr := make([]int, 2*n)
	for i := 0; i < len(arr); i++ {
		arr[i] = v
	}
	return &SegTree{n, arr, v, op}
}

func (seg *SegTree) Update(p int, v int) {
	p += seg.size
	seg.arr[p] = v
	for p > 1 {
		seg.arr[p>>1] = seg.op(seg.arr[p], seg.arr[p^1])
		p >>= 1
	}
}

func (seg *SegTree) Get(l, r int) int {
	res := seg.init_value
	l += seg.size
	r += seg.size
	for l < r {
		if l&1 == 1 {
			res = seg.op(res, seg.arr[l])
			l++
		}
		if r&1 == 1 {
			r--
			res = seg.op(res, seg.arr[r])
		}
		l >>= 1
		r >>= 1
	}
	return res
}
