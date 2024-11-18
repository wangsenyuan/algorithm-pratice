package main

import "testing"

func runSample(t *testing.T, n int, p int, expect int) {
	res := solve(n, p)
	if res != expect {
		t.Errorf("Sample %d %d, expect %d, but got %d", n, p, expect, res)
	}
}
func TestSample1(t *testing.T) {
	runSample(t, 24, 0, 2)
}

func TestSample2(t *testing.T) {
	runSample(t, 24, 1, 3)
}

func TestSample3(t *testing.T) {
	runSample(t, 24, -1, 4)
}

func TestSample4(t *testing.T) {
	runSample(t, 4, -7, 2)
}

func TestSample5(t *testing.T) {
	runSample(t, 1, 1, -1)
}
