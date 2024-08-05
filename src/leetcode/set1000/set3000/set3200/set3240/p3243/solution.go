package p3243

func shortestDistanceAfterQueries(n int, queries [][]int) []int {
	set := NewSegTree(n, n, func(a, b int) int {
		return min(a, b)
	})

	for i := 0; i < n; i++ {
		set.Update(i, i)
	}

	cnt := n

	ans := make([]int, len(queries))

	for i, cur := range queries {
		l, r := cur[0], cur[1]
		for {
			x := set.Get(l+1, r)
			if x == n {
				break
			}
			cnt--
			set.Update(x, n)
		}
		ans[i] = cnt - 1
	}

	return ans
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
