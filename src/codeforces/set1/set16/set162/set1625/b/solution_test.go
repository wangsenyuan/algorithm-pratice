package main

import "testing"

func runSample(t *testing.T, A []int, expect int) {
	res := solve(A)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{3, 1, 5, 2, 1, 3, 4}
	expect := 4
	runSample(t, A, expect)
}

func TestSample2(t *testing.T) {
	A := []int{1, 1, 1, 1, 1, 1}
	expect := 5
	runSample(t, A, expect)
}

func TestSample3(t *testing.T) {
	A := []int{1, 4, 2, 8, 5, 7}
	expect := -1
	runSample(t, A, expect)
}

func TestSample5(t *testing.T) {
	A := []int{1, 1}
	expect := 1
	runSample(t, A, expect)
}
