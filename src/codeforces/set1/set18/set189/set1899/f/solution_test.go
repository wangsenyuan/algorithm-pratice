package main

import "testing"

func runSample(t *testing.T, n int, d []int) {
	_, res := solve(n, d)

	parent := make([]int, n)

	for i := 0; i < n; i++ {
		parent[i] = i - 1
	}
	que := make([]int, n)

	dist := make([]int, n)

	check := func(x int) bool {
		g := NewGraph(n, n)
		for i := 0; i < n; i++ {
			if parent[i] >= 0 {
				g.AddEdge(parent[i], i)
			}
		}
		var front, tail int
		que[front] = 0
		front++
		var ends []int
		for tail < front {
			u := que[tail]
			tail++
			if g.nodes[u] == 0 {
				ends = append(ends, u)
			}
			for i := g.nodes[u]; i > 0; i = g.next[i] {
				v := g.to[i]
				dist[v] = dist[u] + 1
				que[front] = v
				front++
			}
		}
		for i := 0; i < len(ends); i++ {
			if dist[ends[i]] == x {
				return true
			}
		}
		if len(ends) != 2 {
			return false
		}
		u, v := ends[0], ends[1]
		return dist[u]+dist[v]-2 == x
	}

	for i, x := range d {
		cur := res[i]
		if cur[0] != -1 {
			u, v2 := cur[0]-1, cur[2]-1
			parent[u] = v2
		}
		if !check(x) {
			t.Fatalf("Sample result %v not correct", res)
		}
	}

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

func TestSample1(t *testing.T) {
	n := 3
	d := []int{2, 2, 2}
	runSample(t, n, d)
}

func TestSample2(t *testing.T) {
	n := 5
	d := []int{4, 2, 3, 4, 3, 2}
	runSample(t, n, d)
}

func TestSample3(t *testing.T) {
	n := 4
	d := []int{2, 3, 3, 2, 2, 2, 3, 2, 2}
	runSample(t, n, d)
}
