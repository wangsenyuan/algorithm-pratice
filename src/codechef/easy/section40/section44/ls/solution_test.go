package main

import "testing"

func runSample(t *testing.T, V []int, k int, expect int64) {
	res := solve(V, k)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	V := []int{3, 5, 4, 2, 6}
	k := 5
	var expect int64 = 23
	runSample(t, V, k, expect)
}
