package main

import "testing"

func runSample(t *testing.T, k int, expect int) {
	res := solve(k)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 8, 3)
}

func TestSample2(t *testing.T) {
	runSample(t, 4, 2)
}
