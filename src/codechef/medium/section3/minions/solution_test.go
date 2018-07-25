package main

import "testing"

func runSample(t *testing.T, n int, A []int64, B []int64, expect int) {
	res := solve(n, A, B)

	if res != expect {
		t.Errorf("Sample %d %v %v, expect %d, but got %d", n, A, B, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	A := []int64{1, 3, 2}
	B := []int64{4, 3, 1}
	expect := 2
	runSample(t, n, A, B, expect)
}

func TestSample2(t *testing.T) {
	n := 2
	A := []int64{3, 1}
	B := []int64{5, 4}
	expect := 0
	runSample(t, n, A, B, expect)
}

func TestSample3(t *testing.T) {
	n := 12
	A := []int64{10, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12}
	B := []int64{100, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 100000000}
	expect := 11
	runSample(t, n, A, B, expect)
}
