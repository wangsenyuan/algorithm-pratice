package main

import "testing"

func runSample(t *testing.T, A []int64, B []int64, expect int) {
	res := solve(A, B)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int64{4, 0, 4}
	B := []int64{7, 9, 8}
	expect := 9
	runSample(t, A, B, expect)
}

func TestSample2(t *testing.T) {
	A := []int64{0, 3}
	B := []int64{10, 9}
	expect := 32
	runSample(t, A, B, expect)
}

func TestSample3(t *testing.T) {
	A := []int64{0, 3, 0, 2}
	B := []int64{6, 8, 1, 5}
	expect := 22
	runSample(t, A, B, expect)
}
