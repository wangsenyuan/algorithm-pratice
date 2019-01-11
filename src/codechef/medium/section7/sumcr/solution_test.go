package main

import "testing"

func runSample(t *testing.T, k int, A []uint64, expect uint64) {
	var i int
	res := solve(k, func() uint64 {
		x := A[i]
		i += 1
		return x
	})

	if res != expect {
		t.Errorf("Sample %d %v, expect %d, but got %d", k, A, expect, res)
	}
}

func TestSample1(t *testing.T) {
	k := 1
	A := []uint64{5}
	var expect uint64 = 2
	runSample(t, k, A, expect)
}
