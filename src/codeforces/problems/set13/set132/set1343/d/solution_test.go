package main

import "testing"

func runSample(t *testing.T, n int, K int, A []int, expect int) {
	res := solve(n, K, A)

	if res != expect {
		t.Errorf("Sample %d %d %v, expect %d, but got %d", n, K, A, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, k := 4, 2
	A := []int{1, 2, 1, 2}
	expect := 0
	runSample(t, n, k, A, expect)
}

func TestSample2(t *testing.T) {
	n, k := 4, 3
	A := []int{1, 2, 2, 1}
	expect := 1
	runSample(t, n, k, A, expect)
}

func TestSample3(t *testing.T) {
	n, k := 8, 7
	A := []int{6, 1, 1, 7, 6, 3, 4, 6}
	expect := 4
	runSample(t, n, k, A, expect)
}

func TestSample4(t *testing.T) {
	n, k := 6, 6
	A := []int{5, 2, 6, 1, 3, 4}
	expect := 2
	runSample(t, n, k, A, expect)
}
