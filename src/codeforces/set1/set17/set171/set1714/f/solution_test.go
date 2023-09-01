package main

import "testing"

func runSample(t *testing.T, n int, d []int, expect bool) {
	res := solve(n, d)
	if len(res) > 0 != expect {
		t.Fatalf("Sample expect %t, but got %v", expect, res)
	}
	if !expect {
		return
	}

	g := NewGraph(n, 2*n)
	for _, e := range res {
		u, v := e[0], e[1]
		u--
		v--
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}
	que := make([]int, n)

	bfs := func(s int) []int {
		dist := make([]int, n)
		for i := 0; i < n; i++ {
			dist[i] = -1
		}
		dist[s] = 0
		var front, end int
		que[end] = s
		end++
		for front < end {
			u := que[front]
			front++
			for i := g.nodes[u]; i > 0; i = g.next[i] {
				v := g.to[i]
				if dist[v] < 0 {
					dist[v] = dist[u] + 1
					que[end] = v
					end++
				}
			}
		}
		return dist
	}
	a := bfs(0)
	// 12, 23, 31
	if d[0] != a[1] || d[2] != a[2] {
		t.Fatalf("Sample result %v not correct", res)
	}
	b := bfs(1)
	if d[1] != b[2] {
		t.Fatalf("Sample result %v not correct", res)
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
	n := 5
	d := []int{1, 2, 1}
	expect := true
	runSample(t, n, d, expect)
}

func TestSample2(t *testing.T) {
	n := 5
	d := []int{2, 2, 2}
	expect := true
	runSample(t, n, d, expect)
}

func TestSample3(t *testing.T) {
	n := 5
	d := []int{2, 2, 3}
	expect := false
	runSample(t, n, d, expect)
}

func TestSample4(t *testing.T) {
	n := 5
	d := []int{2, 2, 4}
	expect := true
	runSample(t, n, d, expect)
}
