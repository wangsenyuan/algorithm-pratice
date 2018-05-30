package main

import "testing"

func runSample(t *testing.T, n, m int, field [][]int, expect int) {
	res := solve(n, m, field)
	if res != expect {
		t.Errorf("Sample %d %d %v, expect %d, but got %d", n, m, field, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, m := 3, 3
	field := [][]int{
		{1, 2, 2},
		{2, 1, 1},
		{2, 1, 1},
	}
	expect := 2
	runSample(t, n, m, field, expect)
}

func TestSample2(t *testing.T) {
	n, m := 3, 3
	field := [][]int{
		{3, 2, 2},
		{2, 1, 1},
		{2, 1, 1},
	}
	expect := 1
	runSample(t, n, m, field, expect)
}
