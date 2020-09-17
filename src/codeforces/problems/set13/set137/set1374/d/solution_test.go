package main

import "testing"

func runSample(t *testing.T, n, k int, A []int, expect int64) {
	res := solve(n, k, A)

	if res != expect {
		t.Errorf("Sample %d %d %v, expect %d, but got %d", n, k, A, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, k := 4, 3
	A := []int{1, 2, 1, 3}
	var expect int64 = 6
	runSample(t, n, k, A, expect)
}

func TestSample2(t *testing.T) {
	n, k := 10, 6
	A := []int{8, 7, 1, 8, 3, 7, 5, 10, 8, 9}
	var expect int64 = 18
	runSample(t, n, k, A, expect)
}

func TestSample3(t *testing.T) {
	n, k := 5, 10
	A := []int{20, 100, 50, 20, 100500}
	var expect int64 = 0
	runSample(t, n, k, A, expect)
}

func TestSample4(t *testing.T) {
	n, k := 10, 25
	A := []int{24, 24, 24, 24, 24, 24, 24, 24, 24, 24}
	var expect int64 = 227
	runSample(t, n, k, A, expect)
}

func TestSample5(t *testing.T) {
	n, k := 8, 8
	A := []int{1, 2, 3, 4, 5, 6, 7, 8}
	var expect int64 = 8
	runSample(t, n, k, A, expect)
}
