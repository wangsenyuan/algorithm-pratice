package main

import "testing"

func runSample(t *testing.T, s string, expect int64) {
	res := solve(len(s), []byte(s))
	if res != expect {
		t.Errorf("Sample %s, expect %d, but got %d", s, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "rgr", 1)
}

func TestSample2(t *testing.T) {
	runSample(t, "rrr", 0)
}

func TestSample3(t *testing.T) {
	runSample(t, "rgb", 0)
}

func TestSample4(t *testing.T) {
	runSample(t, "rgbr", 2)
}
