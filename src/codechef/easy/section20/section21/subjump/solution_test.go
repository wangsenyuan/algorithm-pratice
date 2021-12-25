package main

import "testing"

func runSample(t *testing.T, A []int, expect int64) {
	res := solve(A)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{6, 1, 3}
	var expect int64 = 10

	runSample(t, A, expect)
}

func TestSample2(t *testing.T) {
	A := []int{3, 1, 10, 4}
	var expect int64 = 4

	runSample(t, A, expect)
}

func TestSample3(t *testing.T) {
	A := []int{7, 9, 1}
	var expect int64 = 20

	runSample(t, A, expect)
}

func TestSample4(t *testing.T) {
	A := []int{1, 5}
	var expect int64 = 0

	runSample(t, A, expect)
}
