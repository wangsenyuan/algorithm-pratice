package main

import "testing"

func runSample(t *testing.T, n int, edges [][]int, s int, mx_cnt int, mn_cnt int) {
	mx_res, a, mn_res, b := solve(n, s, edges)

	if mx_res != mx_cnt || mn_res != mn_cnt {
		t.Fatalf("Sample expect %d %d, but got %d %d", mx_cnt, mn_cnt, mx_res, mn_res)
	}

	c := visit(n, edges, s-1, a)
	d := visit(n, edges, s-1, b)

	if c != mx_cnt || d != mn_cnt {
		t.Fatalf("Sample result gives wrong result %d %d", c, d)
	}
}

func visit(n int, edges [][]int, s int, a string) int {
	g := NewGraph(n, len(edges)*2)
	var j int
	for i, e := range edges {
		x, u, v := e[0], e[1]-1, e[2]-1
		if x == 2 {
			if a[j] == '+' {
				g.AddEdge(u, v, i)
			} else {
				g.AddEdge(v, u, i)
			}
			j++
		} else {
			g.AddEdge(u, v, i)
		}
	}
	que := make([]int, n)
	var front, tail int
	que[front] = s
	front++

	vis := make([]bool, n)
	vis[s] = true
	for tail < front {
		u := que[tail]
		tail++
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if !vis[v] {
				vis[v] = true
				que[front] = v
				front++
			}
		}
	}

	return front
}

func TestSample1(t *testing.T) {
	n := 2
	s := 1
	edges := [][]int{
		{1, 1, 2},
		{2, 2, 1},
	}
	mx := 2
	mn := 2
	runSample(t, n, edges, s, mx, mn)
}

func TestSample2(t *testing.T) {
	n := 6
	s := 3
	edges := [][]int{
		{2, 2, 6},
		{1, 4, 5},
		{2, 3, 4},
		{1, 4, 1},
		{1, 3, 1},
		{2, 2, 3},
	}
	mx := 6
	mn := 2
	runSample(t, n, edges, s, mx, mn)
}
