package main

import "testing"

func runSample(t *testing.T, n int, m int, relation [][]int, expect bool) {
	res := solve(n, m, relation)

	if res != expect {
		t.Errorf("Sample %d %d %v, expect %t, but got %t", n, m, relation, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, m := 3, 2
	relation := [][]int{
		{1, 2},
		{2, 3},
	}
	expect := true
	runSample(t, n, m, relation, expect)
}

func TestSample2(t *testing.T) {
	n, m := 3, 3
	relation := [][]int{
		{1, 2},
		{2, 3},
		{1, 3},
	}
	expect := false
	runSample(t, n, m, relation, expect)
}
