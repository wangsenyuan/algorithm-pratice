package main

import "testing"

func runSample(t *testing.T, a, b, c int, expect int) {
	res := solve(a, b, c)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 3, 5, 7, 5)
}

func TestSample2(t *testing.T) {
	runSample(t, 4, 8, 4, 6)
}

func TestSample3(t *testing.T) {
	runSample(t, 13, 10, 13, 13)
}

func TestSample4(t *testing.T) {
	runSample(t, 325, 226, 999, 517)
}

func TestSample5(t *testing.T) {
	runSample(t, 4, 1, 2, 3)
}
