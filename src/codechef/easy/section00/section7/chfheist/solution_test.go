package main

import "testing"

func runSample(t *testing.T, D, d, P, Q int, expect int64) {
	res := solve(D, d, P, Q)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	D, d, P, Q := 2, 1, 1, 1
	var expect int64 = 3
	runSample(t, D, d, P, Q, expect)
}

func TestSample2(t *testing.T) {
	D, d, P, Q := 3, 2, 1, 1
	var expect int64 = 4
	runSample(t, D, d, P, Q, expect)
}

func TestSample3(t *testing.T) {
	D, d, P, Q := 5, 2, 1, 2
	var expect int64 = 13
	runSample(t, D, d, P, Q, expect)
}
