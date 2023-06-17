package main

import "testing"

func runSample(t *testing.T, A []int, expect int) {
	res := solve(A)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{1, 1, 1}
	expect := 0
	runSample(t, A, expect)
}

func TestSample2(t *testing.T) {
	A := []int{1, 2, 3, 4, 5}
	expect := 1
	runSample(t, A, expect)
}

func TestSample3(t *testing.T) {
	A := []int{1, 2, 3, 2, 1}
	expect := 1
	runSample(t, A, expect)
}

func TestSample4(t *testing.T) {
	A := []int{1, 2, 3, 1, 2, 3, 1}
	expect := 2
	runSample(t, A, expect)
}

func TestSample5(t *testing.T) {
	A := []int{2, 2, 1, 2, 3, 2, 1, 2, 3, 1, 2}
	expect := 3
	runSample(t, A, expect)
}
