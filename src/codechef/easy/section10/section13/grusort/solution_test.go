package main

import "testing"

func runSample(t *testing.T, n int, k int, A []int, expect int64) {
	res := solve(n, k, A)

	if res != expect {
		t.Errorf("Sample %d %d %v, expect %d, but got %d", n, k, A, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, k := 3, 2
	A := []int{2, 1, 2}
	expect := int64(666666672)
	runSample(t, n, k, A, expect)
}
