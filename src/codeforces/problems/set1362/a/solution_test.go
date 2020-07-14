package main

import "testing"

func runSample(t *testing.T, a, b uint64, expect int) {
	res := solve(a, b)

	if res != expect {
		t.Errorf("Sample %d %d, expect %d, but got %d", a, b, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 10, 5, 1)
}

func TestSample2(t *testing.T) {
	runSample(t, 11, 44, 1)
}

func TestSample3(t *testing.T) {
	runSample(t, 17, 21, -1)
}

func TestSample4(t *testing.T) {
	runSample(t, 1, 1, 0)
}

func TestSample5(t *testing.T) {
	runSample(t, 96, 3, 2)
}

func TestSample6(t *testing.T) {
	runSample(t, 2, 128, 2)
}

func TestSample7(t *testing.T) {
	runSample(t, 1001, 1100611139403776, 14)
}

func TestSample8(t *testing.T) {
	runSample(t, 1000000000000000000, 1000000000000000000, 0)
}
