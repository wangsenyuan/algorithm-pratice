package main

import "testing"

func runSample(t *testing.T, N int, expect int) {
	res := solve(N)

	if res != expect {
		t.Errorf("Sample %d, expect %d, but got %d", N, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 3, 0)
}

func TestSample2(t *testing.T) {
	runSample(t, 5, 1)
}

func TestSample3(t *testing.T) {
	runSample(t, 10, 2)
}

func TestSample4(t *testing.T) {
	runSample(t, 60, 14)
}

func TestSample5(t *testing.T) {
	runSample(t, 100, 24)
}

func TestSample6(t *testing.T) {
	runSample(t, 1024, 253)
}

func TestSample7(t *testing.T) {
	runSample(t, 23456, 5861)
}

func TestSample8(t *testing.T) {
	runSample(t, 8735373, 2183837)
}
