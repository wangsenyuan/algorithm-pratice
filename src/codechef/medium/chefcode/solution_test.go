package main

import "testing"

func runSample(t *testing.T, n int, A []int64, K int64, expect int) {
	res := solve(n, A, K)

	if res != expect {
		t.Errorf("sample %d %v %d, expect %d, but got %d", n, A, K, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	var K int64 = 4
	A := []int64{1, 2, 3}
	runSample(t, n, A, K, 5)
}
