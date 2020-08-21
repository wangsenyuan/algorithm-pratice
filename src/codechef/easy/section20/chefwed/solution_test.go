package main

import "testing"

func runSample(t *testing.T, n int, k int, F []int, expect int) {
	res := solve(n, k, F)

	if res != expect {
		t.Errorf("Sample %d %d %v, expect %d, but got %d", n, k, F, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, k := 5, 1
	F := []int{5, 1, 3, 3, 3}
	expect := 3
	runSample(t, n, k, F, expect)
}

func TestSample2(t *testing.T) {
	n, k := 5, 14
	F := []int{1, 4, 2, 4, 4}
	expect := 17
	runSample(t, n, k, F, expect)
}

func TestSample3(t *testing.T) {
	n, k := 5, 2
	F := []int{3, 5, 4, 5, 1}
	expect := 4
	runSample(t, n, k, F, expect)
}
