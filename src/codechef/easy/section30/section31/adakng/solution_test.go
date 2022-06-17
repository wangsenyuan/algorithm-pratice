package main

import "testing"

func runSample(t *testing.T, r, c, k int, expect int) {
	res := solve(r, c, k)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 1, 3, 1, 6)	
}

func TestSample2(t *testing.T) {
	runSample(t, 1, 3, 8, 64)	
}
