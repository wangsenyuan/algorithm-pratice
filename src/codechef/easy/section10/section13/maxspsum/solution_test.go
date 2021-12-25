package main

import "testing"

func runSample(t *testing.T, n, k, s int, A []int, expect int64) {
	res := solve(n, k, s, A)

	if res != expect {
		t.Errorf("Sample %d %d %d %v, expect %d, but bot %d", n, k, s, A, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, k, s := 4, 10, 2
	A := []int{14, 2, 7, 15}
	expect := int64(138)
	runSample(t, n, k, s, A, expect)
}
