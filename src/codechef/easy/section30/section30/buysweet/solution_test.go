package main

import "testing"

func runSample(t *testing.T, R int, A, B []int, expect int) {
	res := solve(R, A, B)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	R := 3
	A := []int{2, 3, 4}
	B := []int{1, 1, 1}
	expect := 2
	runSample(t, R, A, B, expect)
}

func TestSample2(t *testing.T) {
	R := 4
	A := []int{5, 4}
	B := []int{1, 2}
	expect := 1
	runSample(t, R, A, B, expect)
}

func TestSample3(t *testing.T) {
	R := 10
	A := []int{4, 4, 5, 5}
	B := []int{1, 2, 4, 2}
	expect := 7
	runSample(t, R, A, B, expect)
}
