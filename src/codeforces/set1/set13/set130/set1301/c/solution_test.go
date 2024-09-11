package main

import "testing"

func runSample(t *testing.T, n int, m int, expect int) {
	res := solve(n, m)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 3, 1, 4)
}

func TestSample2(t *testing.T) {
	runSample(t, 3, 2, 5)
}

func TestSample3(t *testing.T) {
	runSample(t, 3, 3, 6)
}

func TestSample4(t *testing.T) {
	runSample(t, 4, 0, 0)
}

func TestSample5(t *testing.T) {
	runSample(t, 5, 2, 12)
}

func TestSample6(t *testing.T) {
	// 0100
	runSample(t, 4, 1, 6)
}