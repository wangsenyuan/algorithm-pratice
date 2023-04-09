package p2617

func minimumVisitedCells(grid [][]int) int {
	n := len(grid)
	m := len(grid[0])
	rows := make([]*SegTree, n)
	cols := make([]*SegTree, m)

	for i := 0; i < n; i++ {
		rows[i] = NewSegTree(m)
	}
	for j := 0; j < m; j++ {
		cols[j] = NewSegTree(n)
	}
	rows[n-1].Set(m-1, 0)
	cols[m-1].Set(n-1, 0)
	for i := n - 1; i >= 0; i-- {
		for j := m - 1; j >= 0; j-- {
			var cur int = INF
			v := grid[i][j]
			tmp := rows[i].Get(j+1, min(m, v+j+1))
			if tmp < INF {
				cur = tmp + 1
			}
			tmp = cols[j].Get(i+1, min(n, v+i+1))
			if tmp < INF {
				cur = min(cur, tmp+1)
			}
			if cur < INF {
				rows[i].Set(j, cur)
				cols[j].Set(i, cur)
			}
		}
	}
	res := rows[0].Get(0, 1)
	if res == INF {
		return -1
	}
	return res + 1
}

type SegTree struct {
	arr []int
	sz  int
}

const INF = 1e9

func NewSegTree(n int) *SegTree {
	arr := make([]int, 2*n)
	for i := 1; i < 2*n; i++ {
		arr[i] = INF
	}
	return &SegTree{arr, n}
}

func (t *SegTree) Set(p int, v int) {
	p += t.sz
	t.arr[p] = v
	for p > 1 {
		t.arr[p>>1] = min(t.arr[p], t.arr[p^1])
		p >>= 1
	}
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func (t *SegTree) Get(l int, r int) int {
	l += t.sz
	r += t.sz
	var res int = INF
	for l < r {
		if l&1 == 1 {
			res = min(res, t.arr[l])
			l++
		}
		if r&1 == 1 {
			r--
			res = min(res, t.arr[r])
		}
		l >>= 1
		r >>= 1
	}
	return res
}
