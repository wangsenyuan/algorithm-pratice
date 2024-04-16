package main

import "testing"

func runSample(t *testing.T, a []int, expect [][]int) {
	d, res := solve(a)

	if (len(expect) == 0) != (len(res) == 0) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}

	if len(expect) == 0 {
		// no solution
		return
	}
	x := checkAndGetDiameter(a, expect)
	y := checkAndGetDiameter(a, expect)

	if x != y || y != d {
		t.Fatalf("Sample result %v, not correct", res)
	}
}

func checkAndGetDiameter(a []int, edges [][]int) int {
	n := len(a)
	g := NewGraph(n, len(edges)*2)

	deg := make([]int, n)

	for _, e := range edges {
		u, v := e[0]-1, e[1]-1
		deg[u]++
		deg[v]++
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}
	dist := make([]int, n)

	for i := 0; i < n; i++ {
		if deg[i] > a[i] {
			return -1
		}
	}

	que := make([]int, n)

	bfs := func(x int) int {
		for i := 0; i < n; i++ {
			dist[i] = -1
		}
		var front, tail int
		que[front] = x
		dist[x] = 0
		front++
		for tail < front {
			u := que[tail]
			tail++
			for i := g.nodes[u]; i > 0; i = g.next[i] {
				v := g.to[i]
				if dist[v] < 0 {
					dist[v] = dist[u] + 1
					que[front] = v
					front++
				}
			}
		}
		res := 0
		for i := 0; i < n; i++ {
			if dist[i] > dist[res] {
				res = i
			}
		}
		return res
	}

	far := bfs(0)
	next := bfs(far)

	return dist[next]
}

type Graph struct {
	nodes []int
	next  []int
	to    []int
	cur   int
}

func NewGraph(n int, e int) *Graph {
	nodes := make([]int, n)
	e++
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

func TestSample1(t *testing.T) {
	a := []int{2, 2, 2}
	expect := [][]int{
		{1, 2},
		{2, 3},
	}
	runSample(t, a, expect)
}

func TestSample2(t *testing.T) {
	a := []int{1, 4, 1, 1, 1}
	expect := [][]int{
		{1, 2},
		{3, 2},
		{4, 2},
		{5, 2},
	}
	runSample(t, a, expect)
}

func TestSample3(t *testing.T) {
	a := []int{1, 1, 1}

	runSample(t, a, nil)
}

func TestSample4(t *testing.T) {
	a := []int{1, 2, 1, 1, 2}

	runSample(t, a, nil)
}
