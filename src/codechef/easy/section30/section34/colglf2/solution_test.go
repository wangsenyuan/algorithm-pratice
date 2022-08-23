package main

import "testing"

func runSample(t *testing.T, Q []int, L [][]int, expect int) {
	res := int(solve(Q, L))

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	Q := []int{1, 2}
	L := [][]int{
		{2},
		{3, 4},
	}
	expect := 7
	runSample(t, Q, L, expect)
}
