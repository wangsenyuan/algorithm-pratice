package main

import "testing"

func runSample(t *testing.T, wands []int, N int, min_expect int, max_expect int) {
	min_res, max_res := solve(wands, N)

	if min_res != min_expect || max_res != max_expect {
		t.Errorf("Sample expect %d, %d, but got %d %d", min_expect, max_expect, min_res, max_res)
	}
}

func TestSample1(t *testing.T) {
	N := 7
	wands := []int{5, 2, 3}
	min_expect := 2
	max_expect := 3
	runSample(t, wands, N, min_expect, max_expect)
}
