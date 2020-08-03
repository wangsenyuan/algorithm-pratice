package main

import "testing"

func runSample(t *testing.T, n int, B []int, expect int) {
	res := solve(n, B)

	if res != expect {
		t.Errorf("Sample %d %v, expect %d, but got %d", n, B, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	B := []int{2, 0, 1}
	expect := 4
	runSample(t, n, B, expect)
}
