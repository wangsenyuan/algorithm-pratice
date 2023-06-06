package main

import "testing"

func runSample(t *testing.T, X []int64) {
	res := solve(X)

	expect := bruteForce(X)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample0(t *testing.T) {
	X := []int64{1, 2}

	runSample(t, X)
}

func TestSample1(t *testing.T) {
	X := []int64{1, 7}
	runSample(t, X)
}

func TestSample2(t *testing.T) {
	X := []int64{1, 2, 4}
	runSample(t, X)
}

func TestSample3(t *testing.T) {
	X := []int64{5, 5, 5, 5}
	runSample(t, X)
}

func TestSample4(t *testing.T) {
	X := []int64{6, 2, 2, 1, 0}
	runSample(t, X)
}

func TestSample5(t *testing.T) {
	X := []int64{1, 12, 123, 1234, 12345, 123456}
	runSample(t, X)
}
