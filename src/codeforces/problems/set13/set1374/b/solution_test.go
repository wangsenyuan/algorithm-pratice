package main

import "testing"

func runSample(t *testing.T, n int, expect int) {
	res := solve(n)

	if res != expect {
		t.Errorf("Sample %d expect %d, but got %d", n, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 1, 0)
}

func TestSample2(t *testing.T) {
	runSample(t, 2, -1)
}

func TestSample3(t *testing.T) {
	runSample(t, 3, 2)
}

func TestSample4(t *testing.T) {
	runSample(t, 12, -1)
}

func TestSample5(t *testing.T) {
	runSample(t, 12345, -1)
}

func TestSample6(t *testing.T) {
	runSample(t, 15116544, 12)
}

func TestSample7(t *testing.T) {
	runSample(t, 387420489, 36)
}
