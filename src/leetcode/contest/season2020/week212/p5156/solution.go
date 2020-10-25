package p5156

import (
	"sort"
)

const INF = 1 << 30

var dd = []int{-1, 0, 1, 0, -1}

func matrixRankTransform(matrix [][]int) [][]int {
	m := len(matrix)
	n := len(matrix[0])

	ans := make([][]int, m)
	for i := 0; i < m; i++ {
		ans[i] = make([]int, n)
	}

	set := NewUFSet(n * m)
	row := make([]Pair, n)

	adj := make([][]int, n*m)
	for i := 0; i < n*m; i++ {
		adj[i] = make([]int, 0, 10)
	}

	degree := make([]int, n*m)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			row[j] = Pair{matrix[i][j], i*n + j}
		}
		sort.Sort(Pairs(row))

		for j := 0; j+1 < n; j++ {
			if row[j].first == row[j+1].first {
				a := row[j].second
				b := row[j+1].second
				set.Union(a, b)
			}
		}
	}

	col := make([]Pair, m)

	for j := 0; j < n; j++ {
		for i := 0; i < m; i++ {
			col[i] = Pair{matrix[i][j], i*n + j}
		}
		sort.Sort(Pairs(col))
		for i := 0; i+1 < m; i++ {
			if col[i].first == col[i+1].first {
				set.Union(col[i].second, col[i+1].second)
			}
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			row[j] = Pair{matrix[i][j], i*n + j}
		}
		sort.Sort(Pairs(row))

		for j := 0; j+1 < n; j++ {
			if row[j].first != row[j+1].first {
				a := row[j].second
				b := row[j+1].second
				a, b = set.Find(a), set.Find(b)
				adj[a] = append(adj[a], b)
				degree[b]++
			}
		}
	}

	for j := 0; j < n; j++ {
		for i := 0; i < m; i++ {
			col[i] = Pair{matrix[i][j], i*n + j}
		}
		sort.Sort(Pairs(col))
		for i := 0; i+1 < m; i++ {
			if col[i].first != col[i+1].first {
				a := col[i].second
				b := col[i+1].second
				a, b = set.Find(a), set.Find(b)
				adj[a] = append(adj[a], b)
				degree[b]++
			}
		}
	}

	que := make([]int, n*m)

	var front, end int

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			a := i*n + j
			pa := set.Find(a)
			if pa == a && degree[a] == 0 {
				que[end] = a
				ans[i][j] = 1
				end++
			}
		}
	}

	for front < end {
		cur := que[front]
		front++
		x, y := cur/n, cur%n
		for _, b := range adj[cur] {
			u, v := b/n, b%n
			ans[u][v] = max(ans[u][v], ans[x][y]+1)
			degree[b]--
			if degree[b] == 0 {
				que[end] = b
				end++
			}
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			p := set.Find(i*n + j)
			x, y := p/n, p%n
			ans[i][j] = ans[x][y]
		}
	}

	return ans
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

type Pair struct {
	first, second int
}

type Pairs []Pair

func (ps Pairs) Len() int {
	return len(ps)
}

func (ps Pairs) Less(i, j int) bool {
	return ps[i].first < ps[j].first
}

func (ps Pairs) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}

type UFSet struct {
	arr []int
	cnt []int
}

func NewUFSet(n int) UFSet {
	arr := make([]int, n)
	cnt := make([]int, n)

	for i := 0; i < n; i++ {
		arr[i] = i
		cnt[i] = 1
	}

	return UFSet{arr, cnt}
}

func (uf *UFSet) Find(x int) int {
	if uf.arr[x] != x {
		uf.arr[x] = uf.Find(uf.arr[x])
	}
	return uf.arr[x]
}

func (uf *UFSet) Union(a, b int) bool {
	pa := uf.Find(a)
	pb := uf.Find(b)
	if pa == pb {
		return false
	}

	if uf.cnt[pa] < uf.cnt[pb] {
		pa, pb = pb, pa
	}
	uf.cnt[pa] += uf.cnt[pb]
	uf.arr[pb] = pa
	return true
}
