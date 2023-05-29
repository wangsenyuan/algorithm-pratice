package main

import "testing"

func runSample(t *testing.T, H []int64, expect int64) {
	res := solve(H)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	H := []int64{2, 2, 3, 5, 4, 5}
	var expect int64 = 18
	runSample(t, H, expect)
}
