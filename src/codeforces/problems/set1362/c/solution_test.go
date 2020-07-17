package main

import "testing"

func runSample(t *testing.T, num uint64, expect uint64) {
	res := solve(num)

	if res != expect {
		t.Errorf("Sample %d, expect %d, but got %d", num, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 5, 8)
}

func TestSample2(t *testing.T) {
	runSample(t, 7, 11)
}

func TestSample3(t *testing.T) {
	runSample(t, 11, 19)
}

func TestSample4(t *testing.T) {
	runSample(t, 1, 1)
}

func TestSample5(t *testing.T) {
	runSample(t, 2000000000000, 3999999999987)
}
