package main

import "testing"

func runSample(t *testing.T, x int, expect int64) {
	res := solve(x)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 1, 4)
}

func TestSample2(t *testing.T) {
	runSample(t, 3, 25)
}

func TestSample3(t *testing.T) {
	runSample(t, 3151, 11397376)
}
