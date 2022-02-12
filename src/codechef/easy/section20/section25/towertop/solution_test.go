package main

import "testing"

func runSample(t *testing.T, x int64, m int64, expect int64) {
	res := solve(x, m)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 2, 1, 0)
}

func TestSample2(t *testing.T) {
	runSample(t, 2, 2, 1)
}

func TestSample3(t *testing.T) {
	runSample(t, 2, 3, 2)
}

func TestSample4(t *testing.T) {
	runSample(t, 53, 7, 1)
}
