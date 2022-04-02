package main

import "testing"

func runSample(t *testing.T, A []int, expect int) {
	res := solve(A)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{1, 2, 10, 100}
	expect := 7
	runSample(t, A, expect)
}

func TestSample2(t *testing.T) {
	A := []int{100, 100, 100, 1}
	expect := 1
	runSample(t, A, expect)
}

func TestSample3(t *testing.T) {
	A := []int{5, 1, 1, 1, 8}
	expect := 1
	runSample(t, A, expect)
}

func TestSample4(t *testing.T) {
	A := []int{1, 2, 3, 4, 5, 6}
	expect := 3
	runSample(t, A, expect)
}
