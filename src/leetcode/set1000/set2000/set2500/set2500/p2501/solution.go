package p2501

import "sort"

var dd = []int{-1, 0, 1, 0, -1}

func maxPoints(grid [][]int, queries []int) []int {
	m := len(grid)
	n := len(grid[0])
	// 对于queries[i]
	qs := make([]Pair, len(queries))

	for i := 0; i < len(queries); i++ {
		qs[i] = Pair{queries[i], i}
	}

	sort.Slice(qs, func(i, j int) bool {
		return qs[i].Less(qs[j])
	})

	items := make([]Pair, m*n)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			items[i*n+j] = Pair{grid[i][j], i*n + j}
		}
	}

	sort.Slice(items, func(i, j int) bool {
		return items[i].Less(items[j])
	})

	ans := make([]int, len(qs))

	set := NewSet(m * n)

	for i, j := 0, 0; i < len(qs); i++ {

		for j < len(items) && items[j].first < qs[i].first {
			x, y := items[j].second/n, items[j].second%n
			for k := 0; k < 4; k++ {
				u, v := x+dd[k], y+dd[k+1]
				if u >= 0 && u < m && v >= 0 && v < n && grid[u][v] < qs[i].first {
					set.Union(x*n+y, u*n+v)
				}
			}
			j++
		}

		if grid[0][0] < qs[i].first {
			p0 := set.Find(0)
			ans[qs[i].second] = set.cnt[p0]
		}
	}

	return ans
}

type Pair struct {
	first  int
	second int
}

func (this Pair) Less(that Pair) bool {
	return this.first < that.first
}

type Set struct {
	arr []int
	cnt []int
}

func NewSet(n int) *Set {
	arr := make([]int, n)
	cnt := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i
		cnt[i]++
	}
	return &Set{arr, cnt}
}

func (set *Set) Find(x int) int {
	if set.arr[x] != x {
		set.arr[x] = set.Find(set.arr[x])
	}
	return set.arr[x]
}

func (set *Set) Union(a, b int) bool {
	a = set.Find(a)
	b = set.Find(b)
	if a == b {
		return false
	}
	if set.cnt[a] < set.cnt[b] {
		a, b = b, a
	}
	set.cnt[a] += set.cnt[b]
	set.arr[b] = a
	return true
}
