package main

import "testing"

func runSample(t *testing.T, n int, E [][]int, expect bool) {
	res := solve(n, E)

	if len(res) > 0 != expect {
		t.Errorf("Sample expect %t, but got %s", expect, res)
		return
	}
	if !expect {
		return
	}
	var cnt int
	for i, e := range E {
		u, v := e[0], e[1]
		u--
		v--
		if res[u] == '0' && res[v] == '0' {
			t.Fatalf("Sample result %s, not cover edge %d", res, i)
		}
		if res[u] == '1' && res[v] == '1' {
			cnt++
		}
	}
	if cnt > 1 {
		t.Fatalf("Sample result %s, not a lenient cover", res)
	}
}

func TestSample1(t *testing.T) {
	n := 6
	E := [][]int{
		{6, 5},
		{1, 3},
		{2, 4},
		{3, 4},
		{3, 5},
		{4, 6},
	}
	expect := true
	runSample(t, n, E, expect)
}

func TestSample2(t *testing.T) {
	n := 4
	E := [][]int{
		{1, 2},
		{2, 3},
		{3, 4},
		{1, 4},
		{1, 3},
		{2, 4},
	}
	expect := false
	runSample(t, n, E, expect)
}

func TestSample3(t *testing.T) {
	n := 8
	E := [][]int{
		{1, 3},
		{2, 4},
		{3, 5},
		{4, 6},
		{5, 7},
		{6, 8},
		{1, 2},
		{3, 4},
		{5, 6},
		{7, 8},
		{7, 2},
	}
	expect := true
	runSample(t, n, E, expect)
}

func TestSample4(t *testing.T) {
	n := 4
	E := [][]int{
		{1, 2},
		{2, 3},
		{3, 4},
		{1, 3},
		{2, 4},
	}
	expect := true
	runSample(t, n, E, expect)
}

func TestSample5(t *testing.T) {
	n := 10
	E := [][]int{
		{7, 9},
		{5, 3},
		{3, 4},
		{1, 6},
		{9, 4},
		{1, 4},
		{10, 5},
		{7, 1},
		{9, 2},
		{8, 3},
		{7, 3},
		{10, 9},
		{2, 10},
		{9, 8},
		{3, 2},
		{1, 5},
		{10, 7},
		{9, 5},
		{1, 2},
	}
	expect := true
	runSample(t, n, E, expect)
}
