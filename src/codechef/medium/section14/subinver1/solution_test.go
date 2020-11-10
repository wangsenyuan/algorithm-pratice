package main

import "testing"

func runSample(t *testing.T, n, u int, inst [][]int, expect string) {
	res := solve(n, u, inst)

	if res != expect {
		t.Errorf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, u := 10, 10
	inst := [][]int{
		{9, 10},
		{6, 10},
		{9, 10},
		{1, 8},
		{3, 5},
		{3, 3},
		{3, 4},
		{3, 9},
		{4, 8},
		{7, 7},
	}
	expect := "1111100011"
	runSample(t, n, u, inst, expect)
}
