package main

import "testing"

func runSample(t *testing.T, n int, roads [][]int, expect bool) {
	ok, res := solve(n, roads)

	if ok != expect {
		t.Fatalf("Sample expect %t, but got %t, %s", expect, ok, res)
	}

	if !ok {
		return
	}

	m := len(roads)

	for i := 0; i < m; i++ {
		u, v := roads[i][0], roads[i][1]
		u, v = min(u, v), max(u, v)
		for j := 0; j < i; j++ {
			x, y := roads[j][0], roads[j][1]
			x, y = min(x, y), max(x, y)
			if u < x && x < v && v < y || x < u && u < y && y < v {
				// overlap
				if res[i] == res[j] {
					t.Fatalf("Sample result %s, intersect at %d %d", res, i, j)
				}
			}
		}
	}
}

func TestSample1(t *testing.T) {
	n := 4
	roads := [][]int{
		{1, 3},
		{2, 4},
	}
	expect := true
	runSample(t, n, roads, expect)
}

func TestSample2(t *testing.T) {
	n := 6
	roads := [][]int{
		{1, 3},
		{3, 5},
		{5, 1},
	}
	expect := true
	runSample(t, n, roads, expect)
}
