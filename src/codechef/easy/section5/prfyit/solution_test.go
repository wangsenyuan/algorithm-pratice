package main

import "testing"

func runSample(t *testing.T, s string, expect int) {
	res := solve(s)

	if res != expect {
		t.Errorf("Sample %s, expect %d, but got %d", s, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "010111101", 2)
}

func TestSample2(t *testing.T) {
	runSample(t, "1011100001011101", 3)
}

func TestSample3(t *testing.T) {
	runSample(t, "0110", 0)
}

func TestSample4(t *testing.T) {
	runSample(t, "111111", 0)
}
