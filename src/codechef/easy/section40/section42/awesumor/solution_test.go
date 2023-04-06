package main

import "testing"

func runSample(t *testing.T, n int64, expect int) {
	res := solve(n)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	var n int64 = 7
	expect := 6
	runSample(t, n, expect)
}
