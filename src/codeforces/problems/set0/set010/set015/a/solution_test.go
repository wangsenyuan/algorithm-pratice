package main

import "testing"

func runSample(t *testing.T, n int, w int, houses [][]int, expect int) {
	res := solve(n, w, houses)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, w := 2, 2
	houses := [][]int{
		{0, 4},
		{6, 2},
	}
	expect := 4
	runSample(t, n, w, houses, expect)
}

func TestSample2(t *testing.T) {
	n, w := 2, 2
	houses := [][]int{
		{0, 4},
		{5, 2},
	}
	expect := 3
	runSample(t, n, w, houses, expect)
}

func TestSample3(t *testing.T) {
	n, w := 2, 3
	houses := [][]int{
		{0, 4},
		{5, 2},
	}
	expect := 2
	runSample(t, n, w, houses, expect)
}
