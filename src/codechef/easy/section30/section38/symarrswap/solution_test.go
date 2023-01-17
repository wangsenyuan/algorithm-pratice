package main

import "testing"

func runSample(t *testing.T, A []int, B []int, expect int) {
	res := solve(A, B)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{1, 2}
	B := []int{2, 1}
	expect := 0
	runSample(t, A, B, expect)
}

func TestSample2(t *testing.T) {
	A := []int{1, 5, 3}
	B := []int{2, 3, 1}
	expect := 1
	runSample(t, A, B, expect)
}

func TestSample3(t *testing.T) {
	A := []int{4, 2, 5, 1}
	B := []int{5, 3, 4, 1}
	expect := 3
	runSample(t, A, B, expect)
}

func TestSample4(t *testing.T) {
	A := []int{4}
	B := []int{5}
	expect := 0
	runSample(t, A, B, expect)
}
