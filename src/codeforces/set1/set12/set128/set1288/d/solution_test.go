package main

import "testing"

func runSample(t *testing.T, a [][]int, expect []int) {
	res := solve(a)

	x := get(a, expect)
	y := get(a, res)

	if x != y {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func get(a [][]int, v []int) int {
	i, j := v[0]-1, v[1]-1
	res := 1 << 30
	for k := 0; k < len(a[i]); k++ {
		res = min(res, max(a[i][k], a[j][k]))
	}

	return res
}

func TestSample1(t *testing.T) {
	a := [][]int{
		{5, 0, 3, 1, 2},
		{1, 8, 9, 1, 3},
		{1, 2, 3, 4, 5},
		{9, 1, 0, 3, 7},
		{2, 3, 0, 6, 3},
		{6, 4, 1, 7, 0},
	}
	expect := []int{1, 5}
	runSample(t, a, expect)
}
