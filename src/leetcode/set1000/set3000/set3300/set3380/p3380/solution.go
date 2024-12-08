package p3380

import "sort"

func maxRectangleArea(xCoord []int, yCoord []int) int64 {

	n := len(xCoord)

	xs := make(map[int][]int)

	id := make(map[int]int)

	for i := range n {
		x, y := xCoord[i], yCoord[i]
		xs[x] = append(xs[x], y)
		id[y]++
	}

	var ys []int
	for y := range id {
		ys = append(ys, y)
	}
	sort.Ints(ys)

	for i, y := range ys {
		id[y] = i
	}

	arr := make([]int, 0, len(xs))

	for x := range xs {
		arr = append(arr, x)
	}

	sort.Ints(arr)

	pos := NewSegTree(len(ys))
	res := -1
	for _, x := range arr {
		vs := xs[x]
		sort.Ints(vs)
		m := len(vs)
		for i := 0; i < m; i++ {
			ys[i] = id[vs[i]]
		}

		for i := 0; i+1 < m; i++ {
			y0 := ys[i]
			x0 := pos.Get(y0, y0+1)
			if x0 == -inf {
				continue
			}
			y1 := ys[i+1]
			x1 := pos.Get(y1, y1+1)
			if x1 != x0 {
				continue
			}
			// 还必须保证在这个矩形内部不能有点
			z := pos.Get(y0+1, y1)
			if z < x0 {
				res = max(res, (x-x0)*(vs[i+1]-vs[i]))
			}
		}

		for i := 0; i < m; i++ {
			pos.Update(ys[i], x)
		}
	}

	return int64(res)
}

type SegTree struct {
	arr []int
	sz  int
}

const inf = 1 << 28

func NewSegTree(n int) *SegTree {
	arr := make([]int, 2*n)
	for i := 0; i < len(arr); i++ {
		arr[i] = -inf
	}
	return &SegTree{arr, n}
}

func (t *SegTree) Update(p int, v int) {
	p += t.sz
	t.arr[p] = max(t.arr[p], v)
	for p > 1 {
		t.arr[p>>1] = max(t.arr[p], t.arr[p^1])
		p >>= 1
	}
}

func (t *SegTree) Get(l int, r int) int {
	res := -inf
	l += t.sz
	r += t.sz

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
