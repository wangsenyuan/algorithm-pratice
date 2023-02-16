package main

import "testing"

func runSample(t *testing.T, n int, A []int, expect int64) {
	res := solve(n, A)

	if res != expect {
		t.Errorf("Sample %d, %v, expect %d, but got %d", n, A, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 5
	A := []int{1, 4, 3, 7, 10}
	var expect int64 = 1
	runSample(t, n, A, expect)
}

func TestSample2(t *testing.T) {
	n := 3
	A := []int{1, 1, 1}
	var expect int64 = 3
	runSample(t, n, A, expect)
}

func TestSample3(t *testing.T) {
	n := 4
	A := []int{6, 2, 5, 3}
	var expect int64 = 2
	runSample(t, n, A, expect)
}

func TestSample4(t *testing.T) {
	n := 2
	A := []int{2, 4}
	var expect int64 = 0
	runSample(t, n, A, expect)
}
