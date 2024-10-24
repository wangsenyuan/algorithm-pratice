package main

import "testing"

func runSample(t *testing.T, n int, likes [][]int, expect bool) {
	res := solve(n, likes)

	if len(res) > 0 != expect {
		t.Fatalf("Sample expect %t, but got %v", expect, res)
	}

	if !expect {
		return
	}

	if len(res[0])+len(res[1]) != n {
		t.Fatalf("Sample result %v, not correct", res)
	}

	cat := make([]bool, n+1)

	for _, i := range res[1] {
		cat[i] = true
	}

	jury := make([]bool, n+1)

	for _, i := range res[0] {
		jury[i] = true
	}

	for _, like := range likes {
		u, v := like[0], like[1]
		if jury[u] && cat[v] {
			t.Fatalf("Sample result %v, not correct, as jury %d likes cat %d", res, u, v)
		}
	}
}

func TestSample1(t *testing.T) {
	n := 3
	likes := [][]int{
		{1, 1},
		{2, 2},
		{3, 3},
		{1, 3},
	}
	expect := true
	runSample(t, n, likes, expect)
}

func TestSample2(t *testing.T) {
	n := 3
	likes := [][]int{
		{1, 1},
		{1, 2},
		{1, 3},
		{2, 2},
		{3, 1},
		{3, 2},
		{3, 3},
	}
	expect := true
	runSample(t, n, likes, expect)
}

func TestSample3(t *testing.T) {
	n := 1
	likes := [][]int{
		{1, 1},
	}
	expect := false
	runSample(t, n, likes, expect)
}

func TestSample4(t *testing.T) {
	n := 2
	likes := [][]int{
		{1, 1},
		{1, 2},
		{2, 1},
		{2, 2},
	}
	expect := false
	runSample(t, n, likes, expect)
}
