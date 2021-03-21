package main

import "testing"

func runSample(t *testing.T, A []int, expect int) {
	res := solve(len(A), A)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{1, 2, 3}
	expect := 8
	runSample(t, A, expect)
}

func TestSample2(t *testing.T) {
	A := []int{8}
	expect := 2
	runSample(t, A, expect)
}

func TestSample3(t *testing.T) {
	A := []int{5, 1, 4, 2, 3}
	expect := 24
	runSample(t, A, expect)
}

func TestSample4(t *testing.T) {
	A := []int{81, 3}
	expect := 3
	runSample(t, A, expect)
}
