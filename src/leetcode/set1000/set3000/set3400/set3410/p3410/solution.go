package p3410

import "slices"

func maxSubarraySum(nums []int) int64 {
	x := slices.Max(nums)
	if x <= 0 {
		return int64(x)
	}

	tr := NewTree(nums)
	res := tr.val[0]
	negative := make(map[int][]int)
	for i, num := range nums {
		if num < 0 {
			negative[num] = append(negative[num], i)
		}
	}

	for k, ids := range negative {
		for _, i := range ids {
			tr.Update(i, 0)
		}
		res = max(res, tr.val[0])
		for _, i := range ids {
			tr.Update(i, k)
		}
	}

	return int64(res)
}

type Tree struct {
	val  []int
	pref []int
	suf  []int
	sum  []int
	sz   int
}

func merge(val []int, pref []int, suf []int, sum []int, i int) {
	val[i] = max(val[2*i+1], val[2*i+2], suf[2*i+1]+pref[2*i+2])
	pref[i] = max(pref[2*i+1], sum[2*i+1]+pref[2*i+2])
	suf[i] = max(suf[2*i+2], sum[2*i+2]+suf[2*i+1])
	sum[i] = sum[2*i+1] + sum[2*i+2]
}

func NewTree(arr []int) *Tree {
	n := len(arr)
	val := make([]int, 4*n)
	pref := make([]int, 4*n)
	suf := make([]int, 4*n)
	sum := make([]int, 4*n)

	var build func(i int, l int, r int)
	build = func(i int, l int, r int) {
		if l == r {
			val[i] = arr[l]
			pref[i] = max(0, arr[l])
			suf[i] = max(0, arr[l])
			sum[i] = arr[l]
			return
		}
		mid := (l + r) / 2
		build(2*i+1, l, mid)
		build(2*i+2, mid+1, r)
		merge(val, pref, suf, sum, i)
	}

	build(0, 0, n-1)
	return &Tree{val, pref, suf, sum, n}
}

func (tr *Tree) Update(p int, v int) {
	var loop func(i int, l int, r int)
	loop = func(i int, l int, r int) {
		if l == r {
			tr.val[i] = v
			tr.pref[i] = max(0, v)
			tr.suf[i] = max(0, v)
			tr.sum[i] = v
			return
		}
		mid := (l + r) / 2
		if p <= mid {
			loop(2*i+1, l, mid)
		} else {
			loop(2*i+2, mid+1, r)
		}
		merge(tr.val, tr.pref, tr.suf, tr.sum, i)
	}
	loop(0, 0, tr.sz-1)
}
