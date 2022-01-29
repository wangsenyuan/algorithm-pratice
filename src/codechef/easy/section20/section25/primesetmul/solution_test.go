package main

import "testing"

func runSample(t *testing.T, S []int, M int64, expect int64) {
	res := solve(S, M)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	S := []int{3, 5}
	var M int64 = 10
	var expect int64 = 4
	runSample(t, S, M, expect)
}
