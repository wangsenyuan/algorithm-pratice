package main

import "testing"

func runSample(t *testing.T, n int, x int, expect int) {
	res := solve(n, x)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 10, 2, 4)
}

func TestSample2(t *testing.T) {
	runSample(t, 3, 1, 1)
}

func TestSample3(t *testing.T) {
	runSample(t, 76, 4, 9)
}

func TestSample4(t *testing.T) {
	runSample(t, 100, 99, 0)
}

func TestSample5(t *testing.T) {
	runSample(t, 1000000000, 500000000, 1)
}
