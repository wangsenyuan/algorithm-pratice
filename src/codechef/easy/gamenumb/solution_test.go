package main

import "testing"

func runSample(t *testing.T, n int, A []int64, D []int64, k int, B []int64, expect int64) {
	res := solve(n, A, D, k, B)
	if res != expect {
		t.Errorf("sample %d %v %v %d %v, expect %d, but got %d", n, A, D, k, B, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 4
	A := []int64{1, 2, 3, 4}
	D := []int64{2, 2, 2, 2}
	k := 2
	B := []int64{6, 3}
	runSample(t, n, A, D, k, B, 7)
}
