package main

import "testing"

func runSample(t *testing.T, Y int64, expect int64) {
	res := solve(Y)
	if res != expect {
		t.Errorf("Sample %d, expect %d, but got %d", Y, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 10000000000, 69999300)
}

func TestSample2(t *testing.T) {
	runSample(t, 4, 3)
}

func TestSample3(t *testing.T) {
	runSample(t, 3, 2)
}

func TestSample4(t *testing.T) {
	runSample(t, 2, 1)
}
