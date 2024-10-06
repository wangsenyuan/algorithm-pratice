package p3311

import (
	"slices"
	"sort"
)

func constructGridLayout(n int, edges [][]int) [][]int {
	deg := make([]int, n)
	g := NewGraph(n, 2*len(edges)+1)
	for _, e := range edges {
		u, v := e[0], e[1]
		deg[u]++
		deg[v]++
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}
	mn := slices.Min(deg)
	if mn == 1 {
		// 一条直线
		return solve1(n, deg, g)
	}
	// else mn = 2
	// 先找到两条边
	var corners []int
	for i, v := range deg {
		if v == 2 {
			corners = append(corners, i)
		}
	}
	// len(corners) = 4
	que := make([]int, n)

	bfs := func(x int) []int {
		dist := make([]int, n)
		for i := 0; i < n; i++ {
			dist[i] = -1
		}

		dist[x] = 0

		var head, tail int
		que[head] = x
		head++

		for tail < head {
			u := que[tail]
			tail++
			for i := g.nodes[u]; i > 0; i = g.next[i] {
				v := g.to[i]
				if dist[v] < 0 {
					dist[v] = dist[u] + 1
					que[head] = v
					head++
				}
			}
		}
		return dist
	}

	dist0 := bfs(corners[0])
	// 一层层的放置
	sort.Slice(corners, func(i, j int) bool {
		return dist0[corners[i]] < dist0[corners[j]]
	})
	w := dist0[corners[1]]
	h := dist0[corners[2]]

	res := make([][]int, h+1)
	for i := 0; i <= h; i++ {
		res[i] = make([]int, w+1)
	}

	dist1 := bfs(corners[1])

	for i := 0; i < n; i++ {
		d0 := dist0[i]
		d1 := dist1[i]
		// 假设i的位置是(u, v)
		// 那么有 u + v = d0
		// u + w - v = d1
		u := (d0 + d1 - w) / 2
		v := d0 - u
		res[u][v] = i
	}

	return res
}

func solve1(n int, deg []int, g *Graph) [][]int {
	res := make([][]int, 1)
	res[0] = make([]int, n)
	var first int
	for i, v := range deg {
		if v == 1 {
			first = i
			break
		}
	}
	vis := make([]bool, n)

	for i := 0; i < n; i++ {
		res[0][i] = first
		vis[first] = true
		for j := g.nodes[first]; j > 0; j = g.next[j] {
			second := g.to[j]
			if !vis[second] {
				first = second
				break
			}
		}
	}
	return res
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
