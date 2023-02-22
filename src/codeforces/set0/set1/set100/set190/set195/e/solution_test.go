package main

import "testing"

func runSample(t *testing.T, ops [][]int, expect int) {
	res := solve(ops)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	ops := [][]int{
		{},
		{},
		{2, 1},
		{1, 5, 2, 2},
		{1, 2},
		{3, 4},
	}
	expect := 30
	runSample(t, ops, expect)
}

func TestSample2(t *testing.T) {
	ops := [][]int{
		{},
		{1, 5},
		{},
		{},
		{3, 1, 4, 3},
	}
	expect := 9
	runSample(t, ops, expect)
}
