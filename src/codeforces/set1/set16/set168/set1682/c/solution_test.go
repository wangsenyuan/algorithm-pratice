package main

import "testing"

func runSample(t *testing.T, A []int, expect int) {
	res := solve(A)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{6, 6, 6}
	expect := 1
	runSample(t, A, expect)
}

func TestSample2(t *testing.T) {
	A := []int{2, 5, 4, 5, 2, 4}
	expect := 3
	runSample(t, A, expect)
}

func TestSample3(t *testing.T) {
	A := []int{1, 3, 2, 2}
	expect := 2
	runSample(t, A, expect)
}

func TestSample4(t *testing.T) {
	A := []int{1, 1, 2}
	expect := 2
	runSample(t, A, expect)
}

func TestSample5(t *testing.T) {
	A := []int{1, 2, 2}
	expect := 2
	runSample(t, A, expect)
}
