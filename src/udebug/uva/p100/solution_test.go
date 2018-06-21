package main

import "testing"

func runSample(t *testing.T, a, b int, expect int) {
	res := solve(a, b)

	if res != expect {
		t.Errorf("Sample %d %d, expect %d, but got %d", a, b, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 1, 10, 20)
}

func TestSample2(t *testing.T) {
	runSample(t, 100, 200, 125)
}

func TestSample3(t *testing.T) {
	runSample(t, 201, 210, 89)
}

func TestSample4(t *testing.T) {
	runSample(t, 900, 1000, 174)
}
