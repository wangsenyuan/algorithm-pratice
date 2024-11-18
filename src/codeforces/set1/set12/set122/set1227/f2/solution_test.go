package main

import "testing"

func runSample(t *testing.T, n int, k int, h []int, expect int) {
	res := solve(n, k, h)

	if res != expect {
		t.Errorf("Sample %d %d %v, expect %d, but got %d", n, k, h, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, k := 3, 3
	h := []int{1, 3, 1}
	expect := 9
	runSample(t, n, k, h, expect)
}

func TestSample2(t *testing.T) {
	n, k := 5, 5
	h := []int{1, 1, 4, 2, 2}
	expect := 1000
	runSample(t, n, k, h, expect)
}
