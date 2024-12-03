package main

import "testing"

func runSample(t *testing.T, x int, y int, a int, b int, expect int) {
	res := solve(x, y, a, b)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 10, 12, 2, 5, 3)
}

func TestSample2(t *testing.T) {
	runSample(t, 1, 1, 2, 2, 0)
}

func TestSample3(t *testing.T) {
	runSample(t, 52, 311, 13, 27, 4)
}

func TestSample4(t *testing.T) {
	runSample(t, 1000000000, 1000000000, 1, 1, 1000000000)
}

func TestSample5(t *testing.T) {
	runSample(t, 1000000000, 1, 1, 1000000000, 1)
}

func TestSample6(t *testing.T) {
	runSample(t, 1, 1000000000, 1, 1000000000, 1)
}

func TestSample7(t *testing.T) {
	runSample(t, 1, 2, 1, 1, 1)
}

func TestSample8(t *testing.T) {
	runSample(t, 7, 8, 1, 2, 5)
}

func TestSample9(t *testing.T) {
	runSample(t, 4, 1, 2, 3, 0)
}

func TestSample10(t *testing.T) {
	runSample(t, 1, 1, 1, 3, 0)
}
