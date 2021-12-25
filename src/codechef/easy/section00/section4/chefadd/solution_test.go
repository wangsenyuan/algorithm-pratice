package main

import "testing"

func runSample(t *testing.T, a, b, c uint, expect int64) {
	res := solve(a, b, c)

	if res != expect {
		t.Errorf("Sample %d %d %d, expect %d, but got %d", a, b, c, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 1, 2, 3, 2)
}

func TestSample2(t *testing.T) {
	runSample(t, 369, 428, 797, 56)
}
