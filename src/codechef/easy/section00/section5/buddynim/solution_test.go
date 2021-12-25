package main

import "testing"

func runSample(t *testing.T, A []int, B []int, expect bool) {
	res := solve(A, B)
	if res != expect {
		t.Errorf("Sample %v %v, expect %t, but got %t", A, B, expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{1, 1, 1}
	B := []int{3}
	runSample(t, A, B, true)
}

func TestSample2(t *testing.T) {
	A := []int{7}
	B := []int{1, 2, 4}
	runSample(t, A, B, true)
}

func TestSample3(t *testing.T) {
	A := []int{1}
	B := []int{1}
	runSample(t, A, B, false)
}
