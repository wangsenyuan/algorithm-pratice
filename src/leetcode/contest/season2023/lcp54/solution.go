package lcp54

import "sort"

func minimumCost(cost []int, roads [][]int) int64 {
	n := len(cost)
	if n == 1 {
		return int64(cost[0])
	}
	g := NewGraph(n, len(roads)*2)
	deg := make([]int, n)
	for _, r := range roads {
		a, b := r[0], r[1]
		g.AddEdge(a, b)
		g.AddEdge(b, a)
		deg[a]++
		deg[b]++
	}
	disc := make([]int, n)
	low := make([]int, n)

	for i := 0; i < n; i++ {
		low[i] = -1
	}

	bridge := make([]bool, n)
	vis := make([]bool, n)

	var dfs func(p int, u int)
	var time int
	dfs = func(p int, u int) {
		disc[u] = time
		time++
		low[u] = disc[u]
		vis[u] = true
		var cnt int
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if !vis[v] {
				cnt++
				dfs(u, v)
				if low[v] >= disc[u] && p >= 0 {
					bridge[u] = true
				}
				low[u] = min(low[u], low[v])
			} else if p != v {
				low[u] = min(low[u], disc[v])
			}
		}
		if p < 0 && cnt > 1 {
			bridge[u] = true
		}
	}

	dfs(-1, 0)

	var anti map[int]int

	var dfs2 func(p int, u int) int

	dfs2 = func(p int, u int) int {
		vis[u] = true
		res := u
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p == v || vis[v] {
				continue
			}
			if bridge[v] {
				anti[v]++
				continue
			}
			x := dfs2(u, v)
			if cost[x] < cost[res] {
				res = x
			}
		}
		return res
	}
	for i := 0; i < n; i++ {
		vis[i] = false
	}
	var arr []int
	for i := 0; i < n; i++ {
		if deg[i] == 1 {
			arr = append(arr, i)
			continue
		}
		if bridge[i] || vis[i] {
			continue
		}
		anti = make(map[int]int)
		x := dfs2(-1, i)
		if len(anti) == 1 {
			arr = append(arr, x)
		}
		// else assert len(anti) = 2
	}

	if len(arr) == 0 {
		var x int
		for u := range cost {
			if cost[x] > cost[u] {
				x = u
			}
		}
		return int64(cost[x])
	}

	if len(arr) == 1 {
		return int64(cost[arr[0]])
	}

	sort.Slice(arr, func(i, j int) bool {
		return cost[arr[i]] < cost[arr[j]]
	})
	var sum int64
	for i := 0; i+1 < len(arr); i++ {
		sum += int64(cost[arr[i]])
	}

	return sum
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

type Graph struct {
	nodes []int
	next  []int
	to    []int
	cur   int
}

func NewGraph(n int, e int) *Graph {
	nodes := make([]int, n)
	next := make([]int, e+3)
	to := make([]int, e+3)
	var cur int
	return &Graph{nodes, next, to, cur}
}

func (g *Graph) AddEdge(u, v int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
}
