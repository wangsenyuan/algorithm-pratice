package main

import "testing"

func runSample(t *testing.T, n, k int, A []int, expect int64) {
	res := solve(n, k, A)

	if res != expect {
		t.Errorf("Sample %d %d %v, expect %d, but got %d", n, k, A, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, k := 3, 2
	A := []int{1, 3, 1}
	expect := int64(5)
	runSample(t, n, k, A, expect)
}

func TestSample2(t *testing.T) {
	n, k := 3, 6
	A := []int{3, 3, 3}
	expect := int64(12)
	runSample(t, n, k, A, expect)
}

func TestSample3(t *testing.T) {
	n, k := 5, 6
	A := []int{4, 2, 3, 1, 3}
	expect := int64(15)
	runSample(t, n, k, A, expect)
}
