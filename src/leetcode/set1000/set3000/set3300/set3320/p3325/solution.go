package p3325

func numberOfSubstrings(s string, k int) int {
	freq := NewSegTree(26)

	var res int

	add := func(x byte) {
		v := int(x - 'a')
		freq.Update(v, 1)
	}

	rem := func(x byte) {
		v := int(x - 'a')
		freq.Update(v, -1)
	}

	var l int
	for _, x := range []byte(s) {
		add(x)
		for freq.Get(0, 26) >= k {
			rem(s[l])
			if freq.Get(0, 26) < k {
				add(s[l])
				res += l + 1
				break
			}
			l++
		}
	}

	return res
}

type SegTree struct {
	arr []int
	sz  int
}

func NewSegTree(n int) *SegTree {
	arr := make([]int, 2*n)
	return &SegTree{arr, n}
}

func (tr *SegTree) Update(p int, v int) {
	p += tr.sz
	tr.arr[p] += v

	for p > 1 {
		tr.arr[p>>1] = max(tr.arr[p], tr.arr[p^1])
		p >>= 1
	}
}

func (tr *SegTree) Get(l int, r int) int {
	l += tr.sz
	r += tr.sz
	var res int
	for l < r {
		if l&1 == 1 {
			res = max(res, tr.arr[l])
			l++
		}
		if r&1 == 1 {
			r--
			res = max(res, tr.arr[r])
		}
		l >>= 1
		r >>= 1
	}
	return res
}
