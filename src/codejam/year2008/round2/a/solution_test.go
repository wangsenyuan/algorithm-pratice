package main

import "testing"

func runSample(t *testing.T, M, V int, internal [][]int, leaf []int, expect int) {
	res := solve(M, V, internal, leaf)

	if res != expect {
		t.Errorf("Sample %d %d %v %v, expect %d, but got %d", M, V, internal, leaf, expect, res)
	}
}

func TestSample1(t *testing.T) {
	M, V := 9, 1
	internal := [][]int{
		{1, 0},
		{1, 1},
		{1, 1},
		{0, 0},
	}
	leaf := []int{1, 0, 1, 0, 1}
	expect := 1
	runSample(t, M, V, internal, leaf, expect)
}

func TestSample2(t *testing.T) {
	M, V := 5, 0
	internal := [][]int{
		{1, 1},
		{0, 0},
	}
	leaf := []int{1, 1, 0}
	expect := -1
	runSample(t, M, V, internal, leaf, expect)
}
