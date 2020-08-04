package main

import "testing"

func runSample(t *testing.T, n int, A []int64, expect int) {
	res := solve(n, A)

	if res != expect {
		t.Errorf("Sample %d %v, expect %d, but got %d", n, A, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 2
	A := []int64{2, 3}
	expect := 216
	runSample(t, n, A, expect)
}

func TestSample2(t *testing.T) {
	n := 2
	A := []int64{2, 2}
	expect := 8
	runSample(t, n, A, expect)
}
