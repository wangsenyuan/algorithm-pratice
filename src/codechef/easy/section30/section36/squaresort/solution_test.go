package main

import "testing"

func runSample(t *testing.T, A []int64, expect int) {
	res := solve(A)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int64{1, 3, 9, 9}
	expect := 0
	runSample(t, A, expect)
}

func TestSample2(t *testing.T) {
	A := []int64{1, 9, 3, 9}
	expect := 1
	runSample(t, A, expect)
}

func TestSample3(t *testing.T) {
	A := []int64{4, 2, 1}
	expect := 3
	runSample(t, A, expect)
}
