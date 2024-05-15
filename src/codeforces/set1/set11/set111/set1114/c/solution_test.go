package main

import "testing"

func runSample(t *testing.T, n int, b int, expect int) {
	res := solve(n, b)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 6, 9, 1)
}

func TestSample2(t *testing.T) {
	runSample(t, 38, 11, 3)
}

func TestSample3(t *testing.T) {
	runSample(t, 5, 2, 3)
}

func TestSample4(t *testing.T) {
	runSample(t, 5, 10, 1)
}

func TestSample5(t *testing.T) {
	runSample(t, 1000000000000000000, 1000000000000, 20833333333333332)
}
