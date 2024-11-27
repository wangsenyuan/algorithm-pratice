package main

import "testing"

func runSample(t *testing.T, d int, s int, expect string) {
	res := solve(d, s)
	if res != expect {
		t.Errorf("Sample %d %d, expect %s, but got %s", d, s, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 13, 50, "699998")
}

func TestSample2(t *testing.T) {
	runSample(t, 61, 2, "1000000000000000000000000000001")
}

func TestSample3(t *testing.T) {
	runSample(t, 15, 50, "-1")
}

func TestSample4(t *testing.T) {
	runSample(t, 454, 9, "20430")
}
