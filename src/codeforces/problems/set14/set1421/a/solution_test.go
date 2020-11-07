package main

import "testing"

func runSample(t *testing.T, a, b int, expect int) {
	res := solve(a, b)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 6, 12, 10)
}

func TestSample2(t *testing.T) {
	runSample(t, 4, 9, 13)
}

func TestSample3(t *testing.T) {
	runSample(t, 59, 832, 891)
}

func TestSample4(t *testing.T) {
	runSample(t, 4925, 2912, 6237)
}
