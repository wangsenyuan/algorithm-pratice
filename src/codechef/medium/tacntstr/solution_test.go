package main

import "testing"

func runSample(t *testing.T, s string, expect int64) {
	res := solve([]byte(s))

	if res != expect {
		t.Errorf("Sample %s, expect %d, but got %d", s, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "ZAZ", 25)
}

func TestSample2(t *testing.T) {
	runSample(t, "XYZ", 5)
}

func TestSample3(t *testing.T) {
	runSample(t, "A", 25)
}
