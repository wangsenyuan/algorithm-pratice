package p3373

import "slices"

func maxTargetNodes(edges1 [][]int, edges2 [][]int, k int) []int {
	ans1 := prepare(edges1, k)
	if k > 0 {
		ans2 := prepare(edges2, k-1)
		x := slices.Max(ans2)
		for i := 0; i < len(ans1); i++ {
			ans1[i] += x
		}
	}

	return ans1
}

func prepare(edges [][]int, k int) []int {
	n := len(edges) + 1
	g := NewGraph(n, 2*n)
	for _, e := range edges {
		g.AddEdge(e[0], e[1])
		g.AddEdge(e[1], e[0])
	}

	que := make([]int, n)
	dist := make([]int, n)
	bfs := func(x int) int {
		var head, tail int
		for i := 0; i < n; i++ {
			dist[i] = -1
		}
		dist[x] = 0
		que[head] = x
		head++

		for tail < head {
			u := que[tail]
			tail++
			if dist[u] > k {
				return tail - 1
			}
			for i := g.nodes[u]; i > 0; i = g.next[i] {
				v := g.to[i]
				if dist[v] < 0 {
					dist[v] = dist[u] + 1
					que[head] = v
					head++
				}
			}
		}

		return n
	}

	ans := make([]int, n)

	for i := range n {
		ans[i] = bfs(i)
	}
	return ans
}

type Graph struct {
	nodes []int
	next  []int
	to    []int
	cur   int
}

func NewGraph(n int, e int) *Graph {
	nodes := make([]int, n)
	next := make([]int, e)
	to := make([]int, e)
	return &Graph{nodes, next, to, 0}
}

func (g *Graph) AddEdge(u, v int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
}
