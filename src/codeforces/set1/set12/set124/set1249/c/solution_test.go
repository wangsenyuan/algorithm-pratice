package main

import "testing"

func runSample(t *testing.T, n int, expect int) {
	res := solve(n)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 1, 1)
}

func TestSample2(t *testing.T) {
	runSample(t, 2, 3)
}

func TestSample3(t *testing.T) {
	runSample(t, 6, 9)
}

func TestSample4(t *testing.T) {
	// 1 + 3 + 9
	runSample(t, 13, 13)
}

func TestSample5(t *testing.T) {
	runSample(t, 14, 27)
}

func TestSample6(t *testing.T) {
	// 1, 3, 9, 27, 81, 243, 729,2187,6561
	runSample(t, 3620, 6561)
}

func TestSample7(t *testing.T) {
	// 1 + 3 + 9
	runSample(t, 12, 12)
}

func TestSample8(t *testing.T) {
	// 1 + 3 + 9
	runSample(t, 10000, 19683)
}

func TestSample9(t *testing.T) {
	// 1 + 3 + 9
	runSample(t, 1000000000000000000, 1350851717672992089)
}
