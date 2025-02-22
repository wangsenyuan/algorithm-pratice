package main

import "testing"

func runSample(t *testing.T, A []int, B []int, expect int) {
	res := solve(A, B)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{1, 2, 3, 4, 5}
	B := []int{1, 3, 5, 6}
	expect := 2

	runSample(t, A, B, expect)
}

func TestSample2(t *testing.T) {
	A := []int{1, 3, 5, 2}
	B := []int{2, 3, 4, 5, 6, 1}
	expect := 3

	runSample(t, A, B, expect)
}

func TestSample3(t *testing.T) {
	A := []int{1, 2, 3}
	B := []int{3, 2, 1}
	expect := 2

	runSample(t, A, B, expect)
}

func TestSample4(t *testing.T) {
	A := []int{1}
	B := []int{1, 2}
	expect := 1

	runSample(t, A, B, expect)
}
