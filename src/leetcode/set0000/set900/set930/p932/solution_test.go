package main

import "testing"

func runSample(t *testing.T, N int) {
	res := beautifulArray(N)
	if !isBeautiful(res) {
		t.Errorf("Sample %d, got %v, but is not beautiful", N, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 1)
}

func TestSample2(t *testing.T) {
	runSample(t, 2)
}

func TestSample3(t *testing.T) {
	runSample(t, 3)
}

func TestSample4(t *testing.T) {
	runSample(t, 4)
}

func TestSample5(t *testing.T) {
	runSample(t, 5)
}

func TestSample6(t *testing.T) {
	runSample(t, 6)
}

func TestSample7(t *testing.T) {
	for n := 6; n < 20; n++ {
		runSample(t, n)
	}
}
