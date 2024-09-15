package p3291

const inf = 1 << 30

func minValidStrings(words []string, target string) int {
	m := len(words)
	dp := make([][]int, m)

	for i, word := range words {
		z := zFunction(word + "#" + target)
		dp[i] = z[len(word)+1:]
	}
	n := len(target)

	fp := NewSegTree(n, inf, min)

	for i := n - 1; i >= 0; i-- {
		cur := inf

		for j := 0; j < m; j++ {
			if dp[j][i] == 0 {
				continue
			}
			if dp[j][i] == n-i {
				cur = 1
				break
			}
			x := dp[j][i]
			v := fp.Find(i+1, i+x+1)
			cur = min(cur, v+1)
		}

		fp.Update(i, cur)
	}

	res := fp.Find(0, 1)

	if res >= inf {
		return -1
	}

	return res
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func zFunction(s string) []int {
	n := len(s)
	z := make([]int, n)
	for i, l, r := 1, 0, 0; i < n; i++ {
		if i <= r {
			z[i] = min(r-i+1, z[i-l])
		}
		for i+z[i] < n && s[z[i]] == s[i+z[i]] {
			z[i]++
		}
		if i+z[i]-1 > r {
			l = i
			r = i + z[i] - 1
		}
	}
	return z
}

type SegTree struct {
	arr       []int
	sz        int
	initValue int
	fn        func(int, int) int
}

func NewSegTree(n int, iv int, fn func(int, int) int) *SegTree {
	tr := new(SegTree)
	tr.arr = make([]int, 2*n)
	for i := 0; i < 2*n; i++ {
		tr.arr[i] = iv
	}
	tr.sz = n
	tr.initValue = iv
	tr.fn = fn
	return tr
}

func (tree *SegTree) Update(pos int, v int) {
	pos += tree.sz
	tree.arr[pos] = v
	for pos > 0 {
		tree.arr[pos>>1] = tree.fn(tree.arr[pos], tree.arr[pos^1])
		pos >>= 1
	}
}

func (tree *SegTree) Find(l, r int) int {
	l += tree.sz
	r += tree.sz

	ans := tree.initValue

	for l < r {
		if l&1 == 1 {
			ans = tree.fn(ans, tree.arr[l])
			l++
		}
		if r&1 == 1 {
			r--
			ans = tree.fn(ans, tree.arr[r])
		}
		l >>= 1
		r >>= 1
	}
	return ans
}
