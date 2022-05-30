package main

import "testing"

func runSample(t *testing.T, A []int64, expect int) {
	res := solve(A)

	if res != int64(expect) {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int64{29, 19, 0, 3}
	expect := 1
	runSample(t, A, expect)
}
