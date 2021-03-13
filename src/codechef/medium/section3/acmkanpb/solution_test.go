package main

import "testing"

func runSample(t *testing.T, n int, L int, E [][]int, expect int) {
	res, _ := solve(n, L, E)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 4
	L := 100
	E := [][]int{
		{1, 2, 10},
		{2, 3, 20},
		{1, 3, 40},
		{2, 4, 5},
	}
	expect := 460
	runSample(t, n, L, E, expect)
}
func TestSample2(t *testing.T) {
	n := 3
	L := 100
	E := [][]int{
		{1, 2, 10},
		{2, 3, 10},
		{3, 1, 10},
	}
	expect := -1
	runSample(t, n, L, E, expect)
}
