package main

import "testing"

func runSample(t *testing.T, n, k int, A []int, expect int64) {
	res := solve(n, k, A)
	if res != expect {
		t.Errorf("Sample %d %d %v, expect %d, but got %d", n, k, A, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, k := 5, 2
	A := []int{1, 2, 3, 4, 6}
	expect := int64(11)
	runSample(t, n, k, A, expect)
}

func TestSample2(t *testing.T) {
	n, k := 1, 1
	A := []int{3}
	expect := int64(3)
	runSample(t, n, k, A, expect)
}
func TestSample3(t *testing.T) {
	n, k := 2, 1
	A := []int{2, 2}
	expect := int64(2)
	runSample(t, n, k, A, expect)
}
