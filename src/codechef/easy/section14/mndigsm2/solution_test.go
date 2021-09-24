package main

import "testing"

func runSample(t *testing.T, n, r int64, expect int64) {
	res := solve(n, r)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	var n, r int64 = 216, 7
	var expect int64 = 6
	runSample(t, n, r, expect)
}

func TestSample2(t *testing.T) {
	var n, r int64 = 256, 4
	var expect int64 = 2
	runSample(t, n, r, expect)
}

func TestSample3(t *testing.T) {
	var n, r int64 = 31, 5
	var expect int64 = 3
	runSample(t, n, r, expect)
}
