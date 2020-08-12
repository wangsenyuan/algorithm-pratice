package main

import "testing"

func runSample(t *testing.T, A, B []int, expect int) {
	res := solve(len(A), A, len(B), B)

	if res != expect {
		t.Errorf("Sample %v, %v, expect %d, but got %d", A, B, expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{12, 10, 20, 20, 25, 30}
	B := []int{10, 20, 30}
	expect := 2
	runSample(t, A, B, expect)
}

func TestSample2(t *testing.T) {
	A := []int{1, 2, 2, 2, 2, 2, 2, 2}
	B := []int{1, 2}
	expect := 7
	runSample(t, A, B, expect)
}

func TestSample3(t *testing.T) {
	A := []int{1, 2}
	B := []int{1, 2, 3}
	expect := 0
	runSample(t, A, B, expect)
}

func TestSample4(t *testing.T) {
	A := []int{2}
	B := []int{1}
	expect := 0
	runSample(t, A, B, expect)
}
