package main

import "testing"

func runSample(t *testing.T, n, k int, expect int) {
	res := solve(n, k)

	if res != expect {
		t.Errorf("Sample %d %d, expect %d, but got %d", n, k, expect, res)
	}
}

func TestSample1(t *testing.T) {
		runSample(t, 8, 7, 2)
}

func TestSample2(t *testing.T) {
	runSample(t, 8, 1, 8)
}

func TestSample3(t *testing.T) {
	runSample(t, 999999733, 999999732, 999999733)
}
