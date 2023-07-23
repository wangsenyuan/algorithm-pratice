package p2791

func countPalindromePaths(parent []int, s string) int64 {
	n := len(parent)

	g := NewGraph(n, n)

	for i := 1; i < n; i++ {
		x := int(s[i] - 'a')
		g.AddEdge(parent[i], i, x)
	}

	var ans int64

	cnt := make(map[int]int)
	cnt[0] = 1

	var dfs func(u int, x int)

	dfs = func(u int, x int) {
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			nx := x ^ (1 << g.val[i])
			ans += int64(cnt[nx])
			for j := 0; j < 26; j++ {
				ans += int64(cnt[nx^(1<<j)])
			}
			cnt[nx]++
			dfs(v, nx)
		}

	}

	dfs(0, 0)

	return ans
}

type Graph struct {
	nodes []int
	next  []int
	to    []int
	val   []int
	cur   int
}

func NewGraph(n int, e int) *Graph {
	nodes := make([]int, n)
	next := make([]int, e+1)
	to := make([]int, e+1)
	val := make([]int, e+1)
	return &Graph{nodes, next, to, val, 0}
}

func (g *Graph) AddEdge(u, v, w int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
	g.val[g.cur] = w
}
