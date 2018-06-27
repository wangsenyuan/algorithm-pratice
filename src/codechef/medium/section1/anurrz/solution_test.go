package main

import "testing"

func runSample(t *testing.T, n int, B []int, expect int64) {
	res := solve(n, B)

	if res != expect {
		t.Errorf("Sample %d %v, expect %d, but got %d", n, B, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 2
	B := []int{2, 2}
	runSample(t, n, B, 2)
}

func TestSample2(t *testing.T) {
	n := 3
	B := []int{1, 1, 1}
	runSample(t, n, B, 1)
}
