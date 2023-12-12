package main

import "testing"

func runSample(t *testing.T, n int, E [][]int, expect bool) {
	res := solve(n, E)
	if len(res) > 0 != expect {
		t.Fatalf("Sample expect %t, but got %v", expect, res)
	}

	if !expect {
		return
	}
	deg := make([]int, n)
	for _, e := range res {
		u, v := e[0]-1, e[1]-1
		deg[u]++
		deg[v]++
	}

	var special []int
	var tail []int
	var cycle []int
	for i := 0; i < n; i++ {
		if deg[i] == 0 {
			continue
		}
		if deg[i] == 4 {
			special = append(special, i)
		} else if deg[i] == 2 {
			cycle = append(cycle, i)
		} else if deg[i] == 1 {
			tail = append(tail, i)
		} else {
			t.Fatalf("Sample result %v, not correct", res)
		}
		// reset
		deg[i] = 0
	}
	if len(special) != 1 || len(tail) != 2 {
		t.Fatalf("Sample result %v, not correct", res)
	}

	g := NewGraph(n, 2*n)
	// 剩下的部分，应该是一条线
	for _, e := range res {
		u, v := e[0]-1, e[1]-1
		if u == special[0] || v == special[0] {
			// connect to
			continue
		}
		g.AddEdge(u, v)
		g.AddEdge(v, u)
		deg[u]++
		deg[v]++
	}

	first := -1
	for i := 0; i < n; i++ {
		if deg[i] == 1 {
			first = i
			break
		}
	}

	var dfs func(p int, u int) int
	dfs = func(p int, u int) int {
		sz := 1
		var cnt int
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				if cnt == 1 {
					// 分叉了
					return -1
				}
				tmp := dfs(u, v)
				if tmp < 0 {
					return -1
				}
				sz += tmp
				cnt++
			}
		}
		return sz
	}

	cnt := dfs(-1, first)

	if cnt != len(cycle) {
		t.Fatalf("Sample result %v, not correct", res)
	}
}

func TestSample1(t *testing.T) {
	n := 7
	E := [][]int{
		{1, 2},
		{2, 3},
		{3, 4},
		{4, 1},
		{4, 5},
		{4, 6},
		{4, 2},
		{6, 7},
	}
	expect := true
	runSample(t, n, E, expect)
}

func TestSample2(t *testing.T) {
	n := 7
	E := [][]int{
		{6, 7},
		{1, 2},
		{2, 3},
		{3, 4},
		{4, 1},
		{1, 3},
		{3, 5},
	}
	expect := true
	runSample(t, n, E, expect)
}

func TestSample3(t *testing.T) {
	n := 5
	var E [][]int
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			E = append(E, []int{i + 1, j + 1})
		}
	}
	expect := true
	runSample(t, n, E, expect)
}

func TestSample4(t *testing.T) {
	n := 4
	var E [][]int
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			E = append(E, []int{i + 1, j + 1})
		}
	}
	expect := false
	runSample(t, n, E, expect)
}
