package p3245

func numberOfAlternatingGroups(colors []int, queries [][]int) []int {
	n := len(colors)

	x := NewSegTree(n, n, min)
	y := NewSegTree(n, -1, max)
	sum := NewSegTree(n+1, 0, plus)
	cnt := NewSegTree(n+1, 0, plus)

	prev := func(i int) int {
		j := y.Get(0, i)
		if j < 0 {
			j = y.Get(i+1, n)
		}
		return j
	}

	next := func(i int) int {
		j := x.Get(i+1, n)
		if j == n {
			j = x.Get(0, i)
		}
		return j
	}

	get := func(i int, j int) int {
		sz := (j - i + n) % n
		if sz == 0 {
			sz = n
		}
		return sz
	}

	update := func(l int, r int, pos int, v int) {
		old := get(l, r)
		sum.Add(old, -v*old)
		cnt.Add(old, -v)
		s1 := get(pos, r)
		sum.Add(s1, v*s1)
		cnt.Add(s1, v)
		s2 := get(l, pos)
		sum.Add(s2, v*s2)
		cnt.Add(s2, v)
	}

	ins := func(i int) {
		l := prev(i)
		r := next(i)
		x.Update(i, i)
		y.Update(i, i)
		if l < 0 && r == n {
			sum.Add(n, n)
			cnt.Add(n, 1)
			return
		}

		update(l, r, i, 1)
	}

	del := func(i int) {
		l := prev(i)
		r := next(i)
		x.Update(i, n)
		y.Update(i, -1)
		if l < 0 && r == n {
			// only one
			cnt.Add(n, -1)
			sum.Add(n, -n)
			return
		}
		update(l, r, i, -1)
	}

	for i := 0; i < n; i++ {
		if colors[i] == colors[(i+1)%n] {
			ins(i)
		}
	}

	query := func(sz int) int {
		if x.Get(0, n) == n {
			return n
		}
		s := sum.Get(sz, n+1)
		c := cnt.Get(sz, n+1)
		return s - c*(sz-1)
	}

	change := func(i int, c int) {
		if colors[i] == c {
			return
		}
		if colors[(i-1+n)%n] == colors[i] {
			del((i - 1 + n) % n)
		}
		if colors[i] == colors[(i+1+n)%n] {
			del(i)
		}
		colors[i] = c
		if colors[(i-1+n)%n] == c {
			ins((i - 1 + n) % n)
		}
		if colors[(i+1+n)%n] == c {
			ins(i)
		}
	}

	var res []int
	for _, cur := range queries {
		if cur[0] == 2 {
			i, x := cur[1], cur[2]
			change(i, x)
		} else {
			sz := cur[1]
			res = append(res, query(sz))
		}
	}
	return res
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
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

func (seg *SegTree) Add(p int, v int) {
	p += seg.size
	seg.arr[p] = seg.op(seg.arr[p], v)
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
