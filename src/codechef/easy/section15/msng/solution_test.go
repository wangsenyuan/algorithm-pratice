package main

import "testing"

func runSample(t *testing.T, n int, ss []string, expect int64) {
	res := solve(n, ss)
	if res != expect {
		t.Errorf("Sample %d %v, expect %d, but got %d", n, ss, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	ss := []string{
		"-1 10000",
		"8 20",
		"-1 16",
	}
	expect := int64(16)
	runSample(t, n, ss, expect)
}

func TestSample2(t *testing.T) {
	n := 3
	ss := []string{
		"-1 10100",
		"-1 5A",
		"-1 1011010",
	}
	expect := int64(90)
	runSample(t, n, ss, expect)
}
