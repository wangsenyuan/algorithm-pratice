package main

import "testing"

func runSample(t *testing.T, k int, expect int) {
	res := solve(k)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 1, 1)
}

func TestSample2(t *testing.T) {
	runSample(t, 3, 2)
}

func TestSample3(t *testing.T) {
	runSample(t, 20, 5)
}

func TestSample4(t *testing.T) {
	runSample(t, 38, 2)
}

func TestSample5(t *testing.T) {
	runSample(t, 56, 0)
}

func TestSample6(t *testing.T) {
	runSample(t, 2132, 8)
}

func TestSample7(t *testing.T) {
	runSample(t, 506, 2)
}

func TestSample8(t *testing.T) {
	runSample(t, 999999999999999999, 4)
}

func TestSample9(t *testing.T) {
	runSample(t, 1000000000000000000, 1)
}
