package main

import "testing"

func runSample(t *testing.T, A []int, expect int) {
	res := solve(A)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{2, 2, 2, 2}
	expect := 0
	runSample(t, A, expect)
}

func TestSample2(t *testing.T) {
	A := []int{3, 4, 3, 4, 4}
	expect := 2
	runSample(t, A, expect)
}

func TestSample3(t *testing.T) {
	A := []int{1, 2, 3, 4, 0}
	expect := 3
	runSample(t, A, expect)
}

func TestSample4(t *testing.T) {
	A := []int{5, 5, 5, 6, 6, 6}
	expect := 3
	runSample(t, A, expect)
}
