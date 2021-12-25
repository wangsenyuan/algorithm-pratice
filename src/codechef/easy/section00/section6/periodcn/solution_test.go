package main

import "testing"

func runSample(t *testing.T, left, right int, expect int) {
	res := solve(left, right)
	if res != expect {
		t.Errorf("Sample %d %d, expect %d, but got %d", left, right, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 3, 3, 1)
}

func TestSample2(t *testing.T) {
	runSample(t, 1, 10, 6)
}
