package main

import "testing"

func runSample(t *testing.T, n, k int, ps []int, expect int64) {
	res := solve(n, k, ps)

	if res != expect {
		t.Errorf("Sample %d %d %v, expect %d, but got %d", n, k, ps, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, k := 3, 1
	ps := []int{1, 2, 3}
	expect := int64(2)
	runSample(t, n, k, ps, expect)
}

func TestSample2(t *testing.T) {
	n, k := 4, 3
	ps := []int{1, 2, 2, 3}
	expect := int64(24)
	runSample(t, n, k, ps, expect)
}
