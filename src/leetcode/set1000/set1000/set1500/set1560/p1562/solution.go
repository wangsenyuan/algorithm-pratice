package p1562

func findLatestStep(arr []int, m int) int {
	n := len(arr)

	if m == n {
		return n
	}

	minTree := NewSegTree(n, n+2)
	maxTree := NewSegTree(n, n+2)

	minTree.Update(n-1, n+1)
	maxTree.Update(0, 0)

	for i := n - 1; i >= 0; i-- {
		p := arr[i]
		a := maxTree.GetMin(0, p)
		a = -a
		b := minTree.GetMin(p, n)

		if p-a-1 == m || b-p-1 == m {
			return i
		}
		minTree.Update(p-1, p)
		maxTree.Update(p-1, -p)
	}

	return -1
}

type SegTree struct {
	arr []int
	sz  int
}

func NewSegTree(n int, defaultValue int) SegTree {
	arr := make([]int, 2*n)

	for i := 0; i < len(arr); i++ {
		arr[i] = defaultValue
	}

	return SegTree{arr, n}
}

func (tree *SegTree) Update(p int, v int) {
	p += tree.sz
	tree.arr[p] = v

	for p > 0 {
		tree.arr[p>>1] = min(tree.arr[p], tree.arr[p^1])
		p >>= 1
	}
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func (tree SegTree) GetMin(l, r int) int {
	l += tree.sz
	r += tree.sz

	var res int = 1 << 30

	for l < r {
		if l&1 == 1 {
			res = min(res, tree.arr[l])
			l++
		}
		if r&1 == 1 {
			r--
			res = min(res, tree.arr[r])
		}
		l >>= 1
		r >>= 1
	}
	return res
}
