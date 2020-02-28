package main

import "testing"

func runSample(t *testing.T, N, Z int, A []int, expect int64) {
	res := solve(N, Z, A)

	if res != expect {
		t.Errorf("Sample %d %d %v, expect %d, but got %d", N, Z, Z, expect, res)
	}
}

func TestSample1(t *testing.T) {
	N, Z := 2, 1
	A := []int{50, 55}
	var expect int64 = 55
	runSample(t, N, Z, A, expect)
}

func TestSample2(t *testing.T) {
	N, Z := 2, 1
	A := []int{40, 100}
	var expect int64 = 80
	runSample(t, N, Z, A, expect)
}
