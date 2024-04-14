package p3114

import "math"

const H = 20

const oo = 1 << 60

func minimumValueSum(nums, andValues []int) int {
	n, m := len(nums), len(andValues)
	type args struct{ i, j, and int }
	memo := map[args]int{}
	var dfs func(int, int, int) int
	dfs = func(i, j, and int) int {
		if m-j > n-i { // 剩余元素不足
			return math.MaxInt / 2
		}
		if j == m { // 分了 m 段
			if i == n {
				return 0
			}
			return math.MaxInt / 2
		}
		and &= nums[i]
		if and < andValues[j] { // 剪枝：无法等于 andValues[j]
			return math.MaxInt / 2
		}
		p := args{i, j, and}
		if res, ok := memo[p]; ok {
			return res
		}
		res := dfs(i+1, j, and)  // 不划分
		if and == andValues[j] { // 划分，nums[i] 是这一段的最后一个数
			res = min(res, dfs(i+1, j+1, -1)+nums[i])
		}
		memo[p] = res
		return res
	}
	ans := dfs(0, 0, -1)
	if ans == math.MaxInt/2 {
		return -1
	}
	return ans
}

func minimumValueSum1(nums []int, andValues []int) int {
	n := len(nums)
	prefs := make([][]int, H)
	for i := 0; i < H; i++ {
		prefs[i] = make([]int, n+1)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < H; j++ {
			prefs[j][i+1] = prefs[j][i] + ((nums[i] >> j) & 1)
		}
	}

	dp := NewSegTree(n+1, oo, min)
	dp.Update(0, 0)

	getAndValue := func(l int, r int) int {
		var res int
		for j := 0; j < H; j++ {
			if prefs[j][r]-prefs[j][l] == r-l {
				res |= 1 << j
			}
		}
		return res
	}

	for j, v := range andValues {
		for r := n; r > j; r-- {
			// 先假设它是无效的
			if nums[r-1]&v != v {
				dp.Update(r, oo)
				continue
			}

			if nums[r-1] != v {
				l1 := search(r, func(i int) bool {
					return getAndValue(i, r) > v
				})
				l2 := search(l1, func(i int) bool {
					return getAndValue(i, r) == v
				})

				if getAndValue(l2, r) != v {
					dp.Update(r, oo)
					continue
				}

				tmp := dp.Get(max(l2, j), l1)
				dp.Update(r, tmp+nums[r-1])
				continue
			}
			// nums[r] == v
			l := search(r, func(i int) bool {
				return getAndValue(i, r) == v
			})
			tmp := dp.Get(max(j, l), r)
			dp.Update(r, tmp+nums[r-1])
		}
	}

	res := dp.Get(n, n+1)
	if res == oo {
		return -1
	}
	return res
}

func search(n int, f func(int) bool) int {
	l, r := 0, n
	for l < r {
		mid := (l + r) / 2
		if f(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return r
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
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
