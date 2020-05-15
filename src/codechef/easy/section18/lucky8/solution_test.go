package main

import "testing"

func runSample(t *testing.T, L, R string, expect int) {
	res := solve(L, R)

	if res != expect {
		t.Errorf("Sample %s %s, expect %d, but got %d", L, R, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "1", "10", 0)
}

func TestSample2(t *testing.T) {
	runSample(t, "1", "100", 1)
}

func TestSample3(t *testing.T) {
	runSample(t, "47", "74", 1)
}

func TestSample4(t *testing.T) {
	runSample(t, "0", "7474174", 12)
}
