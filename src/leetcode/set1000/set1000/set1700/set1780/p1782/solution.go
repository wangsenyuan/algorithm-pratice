package p1782

import "sort"

func countPairs(n int, edges [][]int, queries []int) []int {
	E := make([]map[int]int, n)
	deg := make([]int, n)
	addEdge := func(u, v int) {
		if E[u] == nil {
			E[u] = make(map[int]int)
		}
		E[u][v]++
		deg[u]++
	}

	for _, e := range edges {
		u, v := e[0], e[1]
		u--
		v--
		addEdge(u, v)
		addEdge(v, u)
	}

	g := NewGraph(n, len(edges)*2+10)

	for u := 0; u < n; u++ {
		for v, w := range E[u] {
			g.AddEdge(u, v, w)
		}
	}
	D := make([]Pair, n)
	for i := 0; i < n; i++ {
		D[i] = Pair{deg[i], i}
	}

	sort.Slice(D, func(i, j int) bool {
		return D[i].first < D[j].first
	})

	pos := make([]int, n)
	for i := 0; i < n; i++ {
		pos[D[i].second] = i
	}
	ans := make([]int, len(queries))
	for x := 0; x < len(queries); x++ {
		cur := queries[x]
		var k int
		for j := 0; j < n-1; j++ {
			if k <= j {
				k = j + 1
			}
			for k < n && D[j].first+D[k].first <= cur {
				k++
			}

			for j < k-1 && D[j].first+D[k-1].first > cur {
				k--
			}
			if k == n || D[j].first+D[k].first <= cur {
				continue
			}
			// a + D[k] > cur
			// a can match at most n - k items,
			tmp := n - k
			// but need to exclude some neighore
			u := D[j].second
			for i := g.nodes[u]; i > 0; i = g.next[i] {
				v := g.to[i]
				if pos[v] >= k {
					if deg[u]+deg[v]-g.cnt[i] <= cur {
						tmp--
					}
				}
			}
			ans[x] += tmp
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

type Pair struct {
	first, second int
}

type Graph struct {
	nodes []int
	next  []int
	to    []int
	cnt   []int
	cur   int
}

func NewGraph(n int, e int) *Graph {
	nodes := make([]int, n)
	next := make([]int, e)
	to := make([]int, e)
	cnt := make([]int, e)
	return &Graph{nodes, next, to, cnt, 0}
}

func (g *Graph) AddEdge(u, v, w int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
	g.cnt[g.cur] += w
}
