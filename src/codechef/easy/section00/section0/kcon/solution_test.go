package main

import "testing"

func runSample(t *testing.T, n int, A []int, k int, expect int64) {
	res := solve(n, A, k)
	if res != expect {
		t.Errorf("sample %d %v %d, should give %d, but got %d", n, A, k, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, k := 2, 3
	A := []int{1, 2}
	runSample(t, n, A, k, 9)
}

func TestSample2(t *testing.T) {
	n, k := 3, 2
	A := []int{1, -2, 1}
	runSample(t, n, A, k, 2)
}
