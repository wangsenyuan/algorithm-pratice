package main

import "testing"

func runSample(t *testing.T, H, C, T int, expect int64) {
	res := solve(H, C, T)

	if res != expect {
		t.Errorf("Sample %d %d %d, expect %d, but got %d", H, C, T, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 30, 10, 20, 2)
}

func TestSample2(t *testing.T) {
	runSample(t, 41, 15, 30, 7)
}

func TestSample3(t *testing.T) {
	runSample(t, 438837, 375205, 410506, 9)
}
