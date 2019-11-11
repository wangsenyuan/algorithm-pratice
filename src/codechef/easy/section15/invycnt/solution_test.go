package main

import "testing"

func runSample(t *testing.T, n, k int, arr []int, expect int64) {
	res := solve(n, k, arr)

	if res != expect {
		t.Errorf("Sample %d %d %v, expect %d, but got %d", n, k, arr, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, k := 3, 3
	arr := []int{2, 1, 3}
	expect := int64(12)
	runSample(t, n, k, arr, expect)
}

func TestSample2(t *testing.T) {
	n, k := 4, 100
	arr := []int{99, 2, 1000, 24}
	expect := int64(30000)
	runSample(t, n, k, arr, expect)
}
