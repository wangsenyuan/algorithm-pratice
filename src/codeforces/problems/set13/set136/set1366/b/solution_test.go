package main

import "testing"

func runSample(t *testing.T, n, x, m int, ops [][]int, expect int) {
	res := solve(n, x, m, ops)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, x, m := 6, 4, 3
	ops := [][]int{
		{1, 6},
		{2, 3},
		{5, 5},
	}
	expect := 6
	runSample(t, n, x, m, ops, expect)
}
func TestSample2(t *testing.T) {
	n, x, m := 4, 1, 2
	ops := [][]int{
		{2, 4},
		{1, 2},
	}
	expect := 2
	runSample(t, n, x, m, ops, expect)
}

func TestSample3(t *testing.T) {
	n, x, m := 3, 3, 2
	ops := [][]int{
		{2, 3},
		{1, 2},
	}
	expect := 3
	runSample(t, n, x, m, ops, expect)
}
