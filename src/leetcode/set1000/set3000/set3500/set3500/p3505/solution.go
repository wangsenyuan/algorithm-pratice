package p3505

import "sort"

const inf = 1e18

func minOperations(nums []int, x int, k int) int64 {
	n := len(nums)

	a := make([]int, n)
	copy(a, nums)
	sort.Ints(a)
	cur_pos := make([]int, n)
	for i := range n {
		cur_pos[i] = i
	}
	id := make([]int, n)
	for i, v := range nums {
		j := sort.SearchInts(a, v)
		id[i] = cur_pos[j]
		cur_pos[j]++
	}

	tr := NewTree(n)

	dp := make([][]int, n+1)
	for i := range n + 1 {
		dp[i] = make([]int, k+1)
		for j := range k + 1 {
			dp[i][j] = inf
		}
	}

	dp[0][0] = 0

	for i := range n {
		tr.Update(id[i], nums[i], 1)
		copy(dp[i+1], dp[i])
		if i+1 >= x {
			// 如果以i结束的x长度的，作为一个分组
			u := tr.FindPosition(x / 2)
			v := tr.FindPosition(x - x/2 + 1)
			s1 := tr.GetSum(0, u)
			s2 := tr.GetSum(v, n-1)
			s := s2 - s1
			// s2 就是这段的操作次数
			for j := k; j > 0; j-- {
				dp[i+1][j] = min(dp[i+1][j], dp[i-x+1][j-1]+s)
			}
		}
		if i-x+1 >= 0 {
			tr.Update(id[i-x+1], nums[i-x+1], -1)
		}
	}
	return int64(dp[n][k])
}

type Tree struct {
	sum []int
	cnt []int
	sz  int
}

func NewTree(n int) *Tree {
	return &Tree{
		sum: make([]int, 4*n),
		cnt: make([]int, 4*n),
		sz:  n,
	}
}

func (tr *Tree) Update(p int, v int, sign int) {
	var loop func(i int, l int, r int)
	loop = func(i int, l int, r int) {
		if l == r {
			tr.sum[i] += sign * v
			tr.cnt[i] += sign
			return
		}
		mid := (l + r) / 2
		if p <= mid {
			loop(i*2+1, l, mid)
		} else {
			loop(i*2+2, mid+1, r)
		}
		tr.sum[i] = tr.sum[i*2+1] + tr.sum[i*2+2]
		tr.cnt[i] = tr.cnt[i*2+1] + tr.cnt[i*2+2]
	}
	loop(0, 0, tr.sz-1)
}

func (tr *Tree) FindPosition(k int) int {
	var loop func(i int, l int, r int, k int) int
	loop = func(i int, l int, r int, k int) int {
		if l == r {
			return l
		}
		mid := (l + r) / 2
		if tr.cnt[i*2+1] >= k {
			return loop(i*2+1, l, mid, k)
		}
		return loop(i*2+2, mid+1, r, k-tr.cnt[i*2+1])
	}
	return loop(0, 0, tr.sz-1, k)
}

func (tr *Tree) GetSum(L int, R int) int {
	var loop func(i int, l int, r int) int
	loop = func(i int, l int, r int) int {
		if R < l || r < L {
			return 0
		}
		if L <= l && r <= R {
			return tr.sum[i]
		}
		mid := (l + r) / 2
		return loop(i*2+1, l, mid) + loop(i*2+2, mid+1, r)
	}
	return loop(0, 0, tr.sz-1)
}
