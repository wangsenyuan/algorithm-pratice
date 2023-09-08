package main

import "testing"

func runSample(t *testing.T, x uint64, expect uint64) {
	res := solve(x)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	var x uint64 = 16
	var expect uint64 = 5
	runSample(t, x, expect)
}

func TestSample2(t *testing.T) {
	var x uint64 = 1000000000000000000
	var expect uint64 = 23561347048791096
	runSample(t, x, expect)
}
