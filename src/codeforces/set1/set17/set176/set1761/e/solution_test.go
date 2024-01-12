package main

import "testing"

func runSample(t *testing.T, n int, mat []string, expect int) {
	res := solve(n, mat)

	if len(res) != expect {
		t.Fatalf("Sample expect %d, but got %v", expect, res)
	}

	grid := make([][]int, n)
	for i := 0; i < n; i++ {
		grid[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if mat[i][j] == '1' {
				grid[i][j] = 1
			}
		}
	}

	for i := 0; i < len(res); i++ {
		u := res[i] - 1
		for v := 0; v < n; v++ {
			if grid[u][v] == 1 {
				grid[u][v] = 0
				grid[v][u] = 0
			} else {
				grid[u][v] = 1
				grid[v][u] = 1
			}
		}
	}

	g := NewGraph(n, n*n)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				g.AddEdge(i, j)
			}
		}
	}

	vis := make([]bool, n)

	var dfs func(u int) int

	dfs = func(u int) int {
		vis[u] = true
		sz := 1
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if !vis[v] {
				sz += dfs(v)
			}
		}

		return sz
	}

	sz := dfs(0)

	if sz != n {
		t.Fatalf("Sample result %v, not correct", res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	mat := []string{
		"011",
		"100",
		"100",
	}
	expect := 0
	runSample(t, n, mat, expect)
}

func TestSample2(t *testing.T) {
	n := 3
	mat := []string{
		"000",
		"001",
		"010",
	}
	expect := 1
	runSample(t, n, mat, expect)
}

func TestSample3(t *testing.T) {
	n := 4
	mat := []string{
		"0100",
		"1000",
		"0001",
		"0010",
	}
	expect := 2
	runSample(t, n, mat, expect)
}

func TestSample4(t *testing.T) {
	n := 6
	mat := []string{
		"001100",
		"000011",
		"100100",
		"101000",
		"010001",
		"010010",
	}
	expect := 3
	runSample(t, n, mat, expect)
}
func TestSample5(t *testing.T) {
	n := 7
	mat := []string{
		"0100000",
		"1000000",
		"0001110",
		"0010100",
		"0011000",
		"0010001",
		"0000010",
	}
	expect := 1
	runSample(t, n, mat, expect)
}
