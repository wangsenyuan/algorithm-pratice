package main

import "testing"

func runSample(t *testing.T, n int, A []int, B []int, expect int) {
	res := solve(n, A, B)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{1, 4, 3, 8}
	B := []int{15, 8, 11}
	expect := 7
	runSample(t, len(A), A, B, expect)
}

func TestSample2(t *testing.T) {
	A := []int{4, 8}
	B := []int{10}
	expect := 2
	runSample(t, len(A), A, B, expect)
}

func TestSample3(t *testing.T) {
	A := []int{2, 4}
	B := []int{3}
	expect := 1
	runSample(t, len(A), A, B, expect)
}
