package main

import "testing"

func runSample(t *testing.T, n int, expect int) {
	res := solve(n)
	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 4, 6)
}

func TestSample2(t *testing.T) {
	// 2 * 3 =
	runSample(t, 6, 12)
}

func TestSample3(t *testing.T) {
	// 2 * 2 * 2
	// 2 * 3 * 5
	// 4 * 2
	// pow(2, 3) * pow(3, 1)
	runSample(t, 8, 24)
}

func TestSample4(t *testing.T) {
	runSample(t, 7, 64)
}

func TestSample5(t *testing.T) {
	runSample(t, 91, 2985984)
}
