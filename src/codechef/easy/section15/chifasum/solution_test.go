package main

import "testing"

func runSample(t *testing.T, n int, A []int, expect int64) {
	res := solve(n, A)

	if res != expect {
		t.Errorf("Sample %d %v, expect %d, but got %d", n, A, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 6
	A := []int{1, 2, 1, 1, 3, 1}
	expect := int64(6)
	runSample(t, n, A, expect)
}

func TestSample2(t *testing.T) {
	n := 3
	A := []int{4, 1, 4}
	expect := int64(2)
	runSample(t, n, A, expect)
}
