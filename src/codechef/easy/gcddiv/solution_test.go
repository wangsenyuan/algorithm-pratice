package main

import "testing"

func runSample(t *testing.T, n int, A []int, K int, expect bool) {
	res := solve(n, A, K)
	if res != expect {
		t.Errorf("Sample %d %v %d, expect %t, but got %t", n, A, K, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	K := 6
	A := []int{10, 15, 30}
	expect := true
	runSample(t, n, A, K, expect)
}

func TestSample2(t *testing.T) {
	n := 3
	K := 4
	A := []int{5, 10, 20}
	expect := false
	runSample(t, n, A, K, expect)
}
