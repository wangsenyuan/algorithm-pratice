package main

import "testing"

func runSample(t *testing.T, k int, x int64, expect int) {
	res := solve(k, x)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 4, 6, 3)
}

func TestSample2(t *testing.T) {
	runSample(t, 4, 7, 4)
}

func TestSample3(t *testing.T) {
	runSample(t, 1, 2, 1)
}

func TestSample4(t *testing.T) {
	runSample(t, 3, 7, 4)
}

func TestSample5(t *testing.T) {
	runSample(t, 100, 1, 1)
}

func TestSample7(t *testing.T) {
	runSample(t, 2, 5, 3)
}
func TestSample6(t *testing.T) {
	runSample(t, 1000000000, 923456789987654321, 1608737403)
}

func TestSample8(t *testing.T) {
	runSample(t, 3, 8, 4)
}
