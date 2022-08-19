package main

import "testing"

func runSample(t *testing.T, n, x uint64, expect int64) {
	res := solve(n, x)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 4, 5, 0)
}


func TestSample2(t *testing.T) {
	runSample(t, 8, 4, 4)
}

func TestSample3(t *testing.T) {
	runSample(t, 1, 2, 1)
}
