package main

import "testing"

func runSample(t *testing.T, A []int, B []int, expect int) {
	res := solve(A, B)
	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{1, 2, 1}
	B := []int{2, 1, 1}
	expect := 2
	runSample(t, A, B, expect)
}

func TestSample2(t *testing.T) {
	A := []int{0, 0}
	B := []int{0, 1}
	expect := 1
	runSample(t, A, B, expect)
}

func TestSample3(t *testing.T) {
	A := []int{1, 0}
	B := []int{0, 1}
	expect := 2
	runSample(t, A, B, expect)
}
