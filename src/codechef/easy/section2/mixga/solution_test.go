package main

import "testing"

func runSample(t *testing.T, n int, A []int, K int, expect int) {
	res := solve(n, A, K)
	if res != expect {
		t.Errorf("Sample %d %v %d, expect %d, but got %d", n, A, K, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, K := 2, 1
	A := []int{1, 0}
	expect := 1
	runSample(t, n, A, K, expect)
}

func TestSample2(t *testing.T) {
	n, K := 3, 5
	A := []int{0, 1, 0}
	expect := 2
	runSample(t, n, A, K, expect)
}
