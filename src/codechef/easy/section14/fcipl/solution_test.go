package main

import "testing"

func runSample(t *testing.T, p, q, r int, expect int64) {
	res := solve(p, q, r)
	if res != expect {
		t.Errorf("Sample %d %d %d, expect %d, but got %d", p, q, r, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 2, 1, 4, 0)
}

func TestSample2(t *testing.T) {
	runSample(t, 2, 4, 4, 3)
}

func TestSample3(t *testing.T) {
	runSample(t, 2, 5, 4, 10)
}

func TestSample4(t *testing.T) {
	runSample(t, 2, 8, 4, 609)
}
