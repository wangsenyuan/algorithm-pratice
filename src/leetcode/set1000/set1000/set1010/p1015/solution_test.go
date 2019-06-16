package p1015

import "testing"

func runSample(t *testing.T, N int, expect int) {
	res := numDupDigitsAtMostN(N)
	if res != expect {
		t.Errorf("Sample %d, expect %d, but got %d", N, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 20, 1)
}

func TestSample2(t *testing.T) {
	runSample(t, 100, 10)
}

func TestSample3(t *testing.T) {
	runSample(t, 1000, 262)
}

func TestSample4(t *testing.T) {
	runSample(t, 1543, 557)
}

func TestSample5(t *testing.T) {
	runSample(t, 1, 0)
}
