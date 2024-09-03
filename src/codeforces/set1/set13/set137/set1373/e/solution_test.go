package main

import "testing"

func runSample(t *testing.T, n int, k int, expect int) {
	res := solve(n, k)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 1, 0, 1)
}

func TestSample2(t *testing.T) {
	runSample(t, 1, 1, 0)
}

func TestSample3(t *testing.T) {
	runSample(t, 42, 7, 4)
}

func TestSample4(t *testing.T) {
	runSample(t, 13, 7, -1)
}

func TestSample5(t *testing.T) {
	runSample(t, 99, 1, 599998)
}

func TestSample6(t *testing.T) {
	runSample(t, 99, 0, 99999999999)
}

func TestSample7(t *testing.T) {
	runSample(t, 99, 2, 7997)
}

func TestSample8(t *testing.T) {
	runSample(t, 28, 1, 189)
}