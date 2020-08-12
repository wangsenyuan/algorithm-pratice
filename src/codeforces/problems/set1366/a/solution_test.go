package main

import "testing"

func runSample(t *testing.T, a, b int, expect int) {
	res := solve(a, b)

	if res != expect {
		t.Errorf("Sample %d %d, expect %d, but got %d", a, b, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 4, 4, 2)
}

func TestSample2(t *testing.T) {
	runSample(t, 7, 15, 7)
}

func TestSample3(t *testing.T) {
	runSample(t, 8, 7, 5)
}
