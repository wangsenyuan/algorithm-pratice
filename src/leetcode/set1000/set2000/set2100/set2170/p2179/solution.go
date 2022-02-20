package p2179

func goodTriplets(nums1 []int, nums2 []int) int64 {
	n := len(nums1)
	// n <= 100000
	// suppose x, y, z, p(x) < p(y) < p(z) && q(x) < q(y) < q(z)
	// take y
	tree := NewTree(n)
	pos := make([]int, n)
	for i := 0; i < n; i++ {
		pos[nums1[i]] = i
		tree.Set(i, 1)
	}
	R := make([]int, n)
	for i := 0; i < n; i++ {
		tree.Set(pos[nums2[i]], 0)
		R[i] = min(tree.Get(pos[nums2[i]], n), n-i)
	}

	for i := 0; i < n; i++ {
		tree.Set(i, 1)
	}

	var res int64

	for i := n - 1; i >= 0; i-- {
		tree.Set(pos[nums2[i]], 0)
		L := min(i, tree.Get(0, pos[nums2[i]]))
		res += int64(L) * int64(R[i])
	}
	return res
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

type Tree struct {
	arr []int
	n   int
}

func NewTree(n int) *Tree {
	arr := make([]int, 2*n)
	return &Tree{arr, n}
}

func (t *Tree) Reset() {
	for i := 0; i < len(t.arr); i++ {
		t.arr[i] = 0
	}
}

func (t *Tree) Set(p int, v int) {
	arr := t.arr
	p += t.n
	arr[p] = v
	for p > 0 {
		arr[p>>1] = arr[p] + arr[p^1]
		p >>= 1
	}
}

func (t *Tree) Get(l, r int) int {
	l += t.n
	r += t.n
	arr := t.arr
	var res int
	for l < r {
		if l&1 == 1 {
			res += arr[l]
			l++
		}
		if r&1 == 1 {
			r--
			res += arr[r]
		}
		l >>= 1
		r >>= 1
	}
	return res
}
