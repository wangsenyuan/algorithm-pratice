package main

import "testing"

func runSample(t *testing.T, A []int, expect int) {
	res := solve(A)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{5, 2, 3}
	expect := 1
	runSample(t, A, expect)
}

func TestSample2(t *testing.T) {
	A := []int{6, 9}
	expect := 1
	runSample(t, A, expect)
}

func TestSample3(t *testing.T) {
	A := []int{3, 6, 5, 4, 9}
	expect := 2
	runSample(t, A, expect)
}

func TestSample4(t *testing.T) {
	A := []int{3, 3, 3}
	expect := 0
	runSample(t, A, expect)
}
