package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	reader := bufio.NewReader(strings.NewReader(s))

	w, edges, res := process(reader)

	if res == expect {
		return
	}
	if expect == 0 || res == 0 {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
	// check 答案
	n := len(w)
	g := NewGraph(n, 2*n)
	for _, e := range edges {
		u, v := e[0]-1, e[1]-1
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}

	in := make([]int, n)
	out := make([]int, n)
	var pos int

	var dfs func(p int, u int)
	dfs = func(p int, u int) {
		in[u] = pos
		pos++
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				dfs(u, v)
			}
		}
		out[u] = pos
		pos++
	}

	dfs(-1, 0)
	res--

	for u := 0; u < n; u++ {
		if w[u] > w[res] && (in[u] < in[res] || out[u] > out[res]) {
			// 至少对手要有一次选择
			return
		}
	}
	t.Fatalf("Sample result %d, not correct", res)
}

func TestSample1(t *testing.T) {
	s := `4
2 2 4 3
1 2
1 3
2 4`
	expect := 2
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `5
1 2 3 4 5
1 2
2 3
3 4
4 5`
	expect := 0
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `3
1 2 3
1 2
1 3`
	expect := 2
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `5
3 1 3 4 5
1 2
2 3
3 4
4 5`
	expect := 2
	runSample(t, s, expect)
}
func TestSample5(t *testing.T) {
	s := `10
1 2 3 2 4 3 3 4 4 3
1 4
4 6
7 4
6 9
6 5
7 8
1 2
2 3
2 10`
	expect := 10
	runSample(t, s, expect)
}
