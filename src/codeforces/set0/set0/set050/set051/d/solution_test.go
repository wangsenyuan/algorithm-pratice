package main

import "testing"

func runSample(t *testing.T, A []int64, expect int) {
	res := solve(A)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int64{3, 6, 12, 24}
	expect := 0
	runSample(t, A, expect)
}

func TestSample2(t *testing.T) {
	A := []int64{-8, -16, 24, -32}
	expect := 1
	runSample(t, A, expect)
}

func TestSample3(t *testing.T) {
	A := []int64{0, 1, 2, 3}
	expect := 2
	runSample(t, A, expect)
}

func TestSample4(t *testing.T) {
	A := []int64{1, 0, 0, 0, 0, 0, 1}
	expect := 1
	runSample(t, A, expect)
}

func TestSample5(t *testing.T) {
	A := []int64{288, 48, 8}
	expect := 0
	runSample(t, A, expect)
}
