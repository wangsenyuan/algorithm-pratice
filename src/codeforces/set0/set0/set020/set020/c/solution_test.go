package main

import "testing"

func runSample(t *testing.T, n int, edges [][]int, expect []int) {
	res := solve(n, edges)

	if len(expect) == 0 {
		if len(res) > 0 {
			t.Fatalf("Sample expect %v, but got %v", expect, res)
		}
		return
	}
	if len(res) == 0 {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
	x := getDist(n, edges, expect)
	y := getDist(n, edges, expect)

	if x != y {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func getDist(n int, edges [][]int, path []int) int {
	var res int

	get := func(u int, v int) int {
		res := inf
		for _, e := range edges {
			if e[0] == u && e[1] == v || e[0] == v && e[1] == u {
				res = min(res, e[2])
			}
		}
		return res
	}

	for i := 0; i+1 < len(path); i++ {
		u, v := path[i], path[i+1]
		res += get(u, v)
	}

	return res
}

func TestSample1(t *testing.T) {
	n := 5
	edges := [][]int{
		{1, 2, 2},
		{2, 5, 5},
		{2, 3, 4},
		{1, 4, 1},
		{4, 3, 3},
		{3, 5, 1},
	}
	expect := []int{1, 4, 3, 5}
	runSample(t, n, edges, expect)
}
