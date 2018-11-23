package main

import "testing"

func runSample(t *testing.T, k int64, expect int64) {
	res := solve(k)
	if res != expect {
		t.Errorf("Sample %d, expect %d, but got %d", k, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 2, 9)
}

func TestSample2(t *testing.T) {
	runSample(t, 3, 10)
}

func TestSample3(t *testing.T) {
	runSample(t, 100, 170)
}

func TestCompute1(t *testing.T) {
	res := compute(99)
	if res != 29 {
		t.Errorf("Sample need 28, but got %d", res)
	}
}

func TestCompute2(t *testing.T) {
	res := compute(10)
	if res != 3 {
		t.Errorf("Sample need 3, but got %d", res)
	}
}
