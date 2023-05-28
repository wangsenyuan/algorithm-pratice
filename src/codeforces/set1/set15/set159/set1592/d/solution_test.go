package main

import "testing"

func runSample(t *testing.T, n int, E [][]int, expect int) {
	ans := make([][]int, n)
	for i := 0; i < n; i++ {
		ans[i] = make([]int, n)
	}

	for _, e := range E {
		u, v, w := e[0], e[1], e[2]
		u--
		v--
		ans[u][v] = w
		ans[v][u] = w
	}

	ask := func(nodes []int) int {
		var res int
		for i := 0; i < len(nodes); i++ {
			for j := i + 1; j < len(nodes); j++ {
				u, v := nodes[i]-1, nodes[j]-1
				res = max(res, ans[u][v])
			}
		}
		return res
	}

	res := solve(n, E, ask)

	if ans[res[0]-1][res[1]-1] != expect {
		t.Fatalf("Sample expect %d, but got %v", expect, res)
	}
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func TestSample1(t *testing.T) {
	n := 6
	E := [][]int{
		{1, 2, 10},
		{2, 3, 4},
		{2, 4, 6},
		{1, 5, 2},
		{5, 6, 10},
	}
	expect := 10
	runSample(t, n, E, expect)
}
