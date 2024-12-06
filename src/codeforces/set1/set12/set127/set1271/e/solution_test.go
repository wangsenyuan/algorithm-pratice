package main

import "testing"

func runSample(t *testing.T, n int, k int, expect int) {
	res := solve(n, k)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 11, 3, 5)
}

func TestSample2(t *testing.T) {
	runSample(t, 11, 6, 4)
}

func TestSample3(t *testing.T) {
	runSample(t, 20, 20, 1)
}

func TestSample4(t *testing.T) {
	runSample(t, 14, 5, 6)
}

func TestSample5(t *testing.T) {
	runSample(t, 1000000, 100, 31248)
}

func TestSample6(t *testing.T) {
	runSample(t, 1e8, 2, 99999998)
}

func TestSample7(t *testing.T) {
	runSample(t, 13, 2, 12)
}

func TestSample8(t *testing.T) {
	runSample(t, 766540997167959122, 63301807306884502, 40)
}
