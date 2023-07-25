package main

import "testing"

func runSample(t *testing.T, n int64, s int, expect int64) {
	res := solve(n, s)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	var n int64 = 2
	s := 1
	var expect int64 = 8
	runSample(t, n, s, expect)
}
