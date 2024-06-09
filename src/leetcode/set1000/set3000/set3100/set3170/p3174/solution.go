package p3174

import "sort"

func maximumLength(nums []int, k int) int {
	fs := map[int][]int{}
	records := make([]struct{ mx, mx2, num int }, k+1)
	for _, x := range nums {
		if fs[x] == nil {
			fs[x] = make([]int, k+1)
		}
		f := fs[x]
		for j := k; j >= 0; j-- {
			f[j]++
			if j > 0 {
				p := records[j-1]
				m := p.mx
				if x == p.num {
					m = p.mx2
				}
				f[j] = max(f[j], m+1)
			}

			// records[j] 维护 fs[.][j] 的 mx,mx2,num
			v := f[j]
			p := &records[j]
			if v > p.mx {
				if x != p.num {
					p.num = x
					p.mx2 = p.mx
				}
				p.mx = v
			} else if x != p.num && v > p.mx2 {
				p.mx2 = v
			}
		}
	}
	return records[k].mx
}

func maximumLength1(nums []int, k int) int {
	// dp[i][0] = Pair{x, y}
	// 表示在有i次不一样的情况下，最长的序列是x，且最后一个是y
	n := len(nums)
	arr := make([]int, n)
	copy(arr, nums)
	sort.Ints(arr)
	var m int
	for i := 1; i <= n; i++ {
		if i == n || arr[i] > arr[i-1] {
			arr[m] = arr[i-1]
			m++
		}
	}

	getPosition := func(num int) int {
		return sort.Search(m, func(i int) bool {
			return arr[i] >= num
		})
	}

	pos := make([]int, n)
	for i, num := range nums {
		pos[i] = getPosition(num)
	}

	dp := make([]*SegTree, k+1)

	for i := 0; i <= k; i++ {
		dp[i] = NewSegTree(m)
	}

	for _, i := range pos {
		for j := k; j >= 0; j-- {
			x := dp[j].Find(i, i+1)
			if j > 0 {
				x = max(x, dp[j-1].Find(0, i), dp[j-1].Find(i+1, m))
			}
			x++
			dp[j].Update(i, x)
		}
	}
	return dp[k].Find(0, m)
}

type SegTree struct {
	arr []int
	sz  int
}

func NewSegTree(n int) *SegTree {
	tr := new(SegTree)
	tr.arr = make([]int, 2*n)

	tr.sz = n
	return tr
}

func (tree *SegTree) Update(pos int, v int) {
	pos += tree.sz
	tree.arr[pos] = v
	for pos > 0 {
		tree.arr[pos>>1] = max(tree.arr[pos], tree.arr[pos^1])
		pos >>= 1
	}
}

func (tree *SegTree) Find(l, r int) int {
	l += tree.sz
	r += tree.sz

	ans := 0

	for l < r {
		if l&1 == 1 {
			ans = max(ans, tree.arr[l])
			l++
		}
		if r&1 == 1 {
			r--
			ans = max(ans, tree.arr[r])
		}
		l >>= 1
		r >>= 1
	}
	return ans
}
